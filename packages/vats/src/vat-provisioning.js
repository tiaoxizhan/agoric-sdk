// @ts-check
import { E, Far } from '@endo/far';
import { makeNotifierKit } from '@agoric/notifier';
import { heapZone } from '@agoric/zone';
import {
  makeSyncMethodCallback,
  prepareGuardedAttenuator,
} from '@agoric/internal/src/callback.js';
import {
  makeNameHubKit,
  NameHubIKit,
  prepareMixinMyAddress,
} from './nameHub.js';

// This vat contains the controller-side provisioning service. To enable local
// testing, it is loaded by both the controller and other ag-solo vat machines.

/** @param {import('@agoric/zone').Zone} zone */
const prepareSpecializedNameAdmin = zone => {
  const mixinMyAddress = prepareMixinMyAddress(zone);

  /** @type {import('@agoric/internal/src/callback.js').AttenuatorMaker<import('./types.js').NamesByAddressAdmin>} */
  const specialize = prepareGuardedAttenuator(zone, NameHubIKit.nameAdmin, {
    tag: 'NamesByAddressAdmin',
  });

  const makeOverrideFacet = zone.exoClass(
    'NamesByAddressAdminFacet',
    undefined, // TODO: interface guard. same as nameAdmin?
    /** @param {import('./types.js').NameAdmin} nameAdmin */
    nameAdmin => ({ nameAdmin }),
    {
      /**
       * @param {string} address
       * @param {string[]} [reserved]
       * @returns {Promise<{ nameHub: NameHub, nameAdmin: import('./types.js').MyAddressNameAdmin}>}
       */
      async provideChild(address, reserved) {
        const { nameAdmin } = this.state;
        const child = await nameAdmin.provideChild(address, reserved);
        return {
          nameHub: child.nameHub,
          nameAdmin: mixinMyAddress(child.nameAdmin, address),
        };
      },
      /** @param {string} address */
      async lookupAdmin(address) {
        const { nameAdmin } = this.state;
        // XXX relies on callers not to provide other admins via update()
        // TODO: enforce?

        /** @type { import('./types').MyAddressNameAdmin } */
        // @ts-expect-error cast
        const myAdmin = nameAdmin.lookupAdmin(address);
        return myAdmin;
      },
    },
  );

  /**
   * @param {import('./types.js').NameAdmin} nameAdmin
   */
  const makeMyAddressNameAdmin = nameAdmin => {
    const overrideFacet = makeOverrideFacet(nameAdmin);
    return specialize({
      target: nameAdmin,
      overrides: {
        provideChild: makeSyncMethodCallback(overrideFacet, 'provideChild'),
        lookupAdmin: makeSyncMethodCallback(overrideFacet, 'lookupAdmin'),
      },
    });
  };

  return makeMyAddressNameAdmin;
};

/**
 * @param {unknown} _vatPowers
 * @param {unknown} _vatParameters
 * @param {import('@agoric/vat-data').Baggage} _baggage
 */
export function buildRootObject(_vatPowers, _vatParameters, _baggage) {
  // TODO: make NameHubKit durable
  const { nameHub: namesByAddress, nameAdmin } = makeNameHubKit();

  // TODO: const zone = makeDurableZone(_baggage);
  const zone = heapZone;
  const makeNamesByAddressAdmin = prepareSpecializedNameAdmin(zone);
  const namesByAddressAdmin = makeNamesByAddressAdmin(nameAdmin);

  let bundler;
  let comms;
  let vattp;

  async function register(b, c, v) {
    bundler = b;
    comms = c;
    vattp = v;
  }

  async function pleaseProvision(nickname, address, powerFlags) {
    let clientFacet;
    const fetch = Far('fetch', {
      async getChainBundle() {
        console.warn('getting chain bundle');
        return E(clientFacet).getChainBundle();
      },
      getConfiguration() {
        console.warn('getting configuration');
        return E(clientFacet).getConfiguration();
      },
    });

    // Add a remote and egress for the pubkey.
    const { transmitter, setReceiver } = await E(vattp).addRemote(address);
    await E(comms).addRemote(address, transmitter, setReceiver);

    const INDEX = 1;
    await E(comms).addEgress(address, INDEX, fetch);

    // Do this here so that any side-effects don't happen unless
    // the egress has been successfully added.
    clientFacet = E(bundler)
      .createClientFacet(nickname, address, powerFlags || [])
      .catch(e => {
        console.warn(`Failed to create client facet:`, e);
        // Emulate with existing createUserBundle.
        const chainBundle = E(bundler).createUserBundle(
          nickname,
          address,
          powerFlags || [],
        );
        // Update the notifier when the chainBundle resolves.
        const { notifier, updater } = makeNotifierKit();
        void E.when(chainBundle, clientHome => {
          updater.updateState(harden({ clientHome, clientAddress: address }));
        });
        return Far('emulatedClientFacet', {
          getChainBundle: () => chainBundle,
          getConfiguration: () => notifier,
        });
      });

    return { ingressIndex: INDEX };
  }

  return Far('root', {
    register,
    pleaseProvision,
    getNamesByAddressKit: () => harden({ namesByAddress, namesByAddressAdmin }),
  });
}
