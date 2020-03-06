// eslint-disable-next-line import/no-extraneous-dependencies
import { test } from 'tape-promise/tape';

import { areRightsConserved, transpose } from '../../src/areRightsConserved';
import { setup } from './setupBasicMints';

const makeAmountMatrix = (amountMaths, extentMatrix) =>
  extentMatrix.map(row => row.map((extent, i) => amountMaths[i].make(extent)));

test('transpose', t => {
  try {
    t.deepEquals(
      transpose([
        [1, 2, 3],
        [4, 5, 6],
      ]),
      [
        [1, 4],
        [2, 5],
        [3, 6],
      ],
    );
  } catch (e) {
    t.assert(false, e);
  } finally {
    t.end();
  }
});

// rights are conserved for amount with Nat extents
test(`areRightsConserved - true for amount with nat extents`, t => {
  try {
    const { amountMaths } = setup();
    const oldExtents = [
      [0, 1, 0],
      [4, 1, 0],
      [6, 3, 0],
    ];
    const newExtents = [
      [1, 2, 0],
      [3, 1, 0],
      [6, 2, 0],
    ];

    const oldAmounts = makeAmountMatrix(amountMaths, oldExtents);
    const newAmounts = makeAmountMatrix(amountMaths, newExtents);

    t.ok(areRightsConserved(amountMaths, oldAmounts, newAmounts));
  } catch (e) {
    t.assert(false, e);
  } finally {
    t.end();
  }
});

// rights are *not* conserved for amount with Nat extents
test(`areRightsConserved - false for amount with Nat extents`, t => {
  try {
    const { amountMaths } = setup();
    const oldExtents = [
      [0, 1, 4],
      [4, 1, 0],
      [6, 3, 0],
    ];
    const newExtents = [
      [1, 2, 0],
      [3, 1, 0],
      [6, 2, 0],
    ];

    const oldAmounts = makeAmountMatrix(amountMaths, oldExtents);
    const newAmounts = makeAmountMatrix(amountMaths, newExtents);

    t.notOk(areRightsConserved(amountMaths, oldAmounts, newAmounts));
  } catch (e) {
    t.assert(false, e);
  } finally {
    t.end();
  }
});

test(`areRightsConserved - empty arrays`, t => {
  try {
    const { amountMaths } = setup();
    const oldAmounts = [[], [], []];
    const newAmounts = [[], [], []];

    t.ok(areRightsConserved(amountMaths, oldAmounts, newAmounts));
  } catch (e) {
    t.assert(false, e);
  } finally {
    t.end();
  }
});

// TODO: add tests for non-Nat extents
