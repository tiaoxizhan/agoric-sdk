(this["webpackJsonp@agoric/wallet-frontend"]=this["webpackJsonp@agoric/wallet-frontend"]||[]).push([[0],{54:function(e,t,a){e.exports=a(65)},65:function(e,t,a){"use strict";a.r(t);var n=a(0),r=a.n(n),c=a(12),o=a(49),l=a(96),i=a(15),u=a(33),s=a.n(u),m=a(42);function p(e){return E.apply(this,arguments)}function E(){return(E=Object(m.a)(s.a.mark((function e(t){return s.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",fetch("/vat",{method:"POST",body:JSON.stringify(t),headers:{"Content-Type":"application/json"}}).then((function(e){return e.json()})).then((function(e){var t=e.ok,a=e.res;return t?a:{}})).catch((function(e){console.log("Fetch Error",e)})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}var d=null;function f(e){var t=e.onConnect,a=e.onDisconnect,n=e.onMessage;d=new WebSocket(function(){var e=new URL("http://127.0.0.1:8000");return e.protocol="ws",e}()),t&&d.addEventListener("open",(function(){return t()})),a&&d.addEventListener("close",(function(){return a()})),n&&d.addEventListener("message",(function(e){var t=e.data;return n(t)}))}function g(){return!!d}function b(){g()&&(d.close(),d=null)}var v=function(e){return{type:"UPDATE_PURSES",payload:e}},O=function(e){return{type:"UPDATE_INBOX",payload:e}},y=a(16);var h=function(e,t){var a=t.type,n=t.payload;switch(a){case"ACTIVATE_CONNECTION":return function(e){return Object(y.a)({},e,{active:!0})}(e);case"DEACTIVATE_CONNECTION":return function(e){return Object(y.a)({},e,{active:!1})}(e);case"SERVER_CONNECTED":return function(e){return Object(y.a)({},e,{connected:!0})}(e);case"SERVER_DISCONNECTED":return function(e){return Object(y.a)({},e,{connected:!1})}(e);case"UPDATE_PURSES":return function(e,t){return Object(y.a)({},e,{purses:t})}(e,n);case"UPDATE_INBOX":return function(e,t){return Object(y.a)({},e,{inbox:t})}(e,n);case"REJECT_OFFER":return function(e,t){return p({type:"walletDeclineOffer",data:t}),e}(e,n);case"CANCEL_OFFER":return function(e,t){return p({type:"walletCancelOffer",data:t}),e}(e,n);case"CONFIRM_OFFER":return function(e,t){return p({type:"walletAcceptOffer",data:t}),e}(e,n);default:throw new TypeError("Action not supported ".concat(a))}},C=Object(n.createContext)();function N(){return Object(n.useContext)(C)}var T=a(8),j=a(82),R=a(95),A=a(84),w=a(85),k=a(25),F=a(86),I=a(44),x=a.n(I),D=Object(j.a)((function(e){return{appBar:{position:"relative"},title:{flexGrow:1},icon:{marginRight:e.spacing(1)}}}));function _(e){var t=e.children,a=D();return r.a.createElement(A.a,{position:"absolute",className:a.appBar},r.a.createElement(w.a,null,r.a.createElement(x.a,{className:a.icon}),r.a.createElement(k.a,{variant:"h6",color:"inherit",noWrap:!0,className:a.title},r.a.createElement(F.a,{color:"inherit",target:"_blank",href:"https://agoric.com"},"Agoric Wallet")),t))}var S=a(98),P=a(87),V=a(45),W=a.n(V),L=Object(j.a)((function(e){return{divider:{marginRight:e.spacing(2)}}}));function U(){var e=L(),t=N(),a=t.state,c=t.dispatch,o=a.active,l=a.connected,i=a.account;return Object(n.useEffect)((function(){return c({type:"ACTIVATE_CONNECTION"})}),[c]),r.a.createElement(r.a.Fragment,null,l&&r.a.createElement(S.a,{className:e.divider,label:i||"anonymous",avatar:r.a.createElement(W.a,null)}),o?r.a.createElement(P.a,{variant:"contained",onClick:function(){c({type:"DEACTIVATE_CONNECTION"})}},"Disconnect"):r.a.createElement(P.a,{variant:"contained",onClick:function(){c({type:"ACTIVATE_CONNECTION"})}},"Connect"))}var B=a(67),J=a(93),z=a(88),K=a(99),M=a(89),G=a(90),X=a(46),q=a.n(X),H=Object(j.a)((function(e){return{icon:{minWidth:24,marginRight:e.spacing(2)}}}));function Q(){var e=H(),t=N().state.purses;return r.a.createElement(r.a.Fragment,null,r.a.createElement(k.a,{variant:"h6"},"Purses"),Array.isArray(t)&&t.length>0?r.a.createElement(z.a,null,t.map((function(t){var a=t.pursePetname,n=t.brandRegKey,c=t.issuerPetname,o=t.extent;return r.a.createElement(K.a,{key:a,value:a,divider:!0},r.a.createElement(M.a,{className:e.icon},r.a.createElement(q.a,null)),r.a.createElement(G.a,{primary:a,secondary:r.a.createElement(r.a.Fragment,null,r.a.createElement("b",null,o," ",c)," ",n?r.a.createElement("i",null,"(",n,")"):"")}))}))):r.a.createElement(k.a,{color:"inherit"},"No purses."))}var Y=a(4),Z=a(30),$=a(92),ee=a(31),te=a(91),ae=a(97),ne=a(94),re=a(47),ce=a.n(re),oe=a(34),le=a.n(oe),ie=a(48),ue=a.n(ie),se=Object(j.a)((function(e){return{icon:{margin:e.spacing(1)},buttons:{"& button ~ button":{marginLeft:e.spacing(1)}}}})),me=Object(Y.a)((function(e){return{root:{color:e.palette.getContrastText(Z.a[500]),backgroundColor:Z.a[500],"&:hover":{backgroundColor:Z.a[700]}}}}))(te.a),pe=Object(Y.a)((function(e){return{root:{color:e.palette.getContrastText($.a[500]),backgroundColor:$.a[500],"&:hover":{backgroundColor:$.a[700]}}}}))(te.a),Ee=Object(Y.a)((function(e){return{root:{color:e.palette.getContrastText(ee.a[500]),backgroundColor:ee.a[500],"&:hover":{backgroundColor:ee.a[700]}}}}))(te.a),de=Object(Y.a)((function(e){return{root:{width:e.spacing(12),color:Z.a[800],borderColor:Z.a[800]}}}))(S.a),fe=Object(Y.a)((function(e){return{root:{width:e.spacing(12),color:ee.a[800],borderColor:ee.a[800]}}}))(S.a),ge=function(e){return e||"???"};function be(){var e=se(),t=N(),a=t.state,n=t.dispatch,c=a.inbox;return r.a.createElement(r.a.Fragment,null,r.a.createElement(k.a,{variant:"h6"},"Transactions"),Array.isArray(c)&&c.length>0?r.a.createElement(z.a,null,c.reverse().map((function(t){var a=t.requestContext,c=(a=void 0===a?{}:a).date,o=a.origin,l=void 0===o?"unknown origin":o,u=t.id,s=t.instanceRegKey,m=t.instancePetname,p=t.offerRulesTemplate,E=(p=void 0===p?{}:p).offer,d=void 0===E?{}:E,f=p.want,g=void 0===f?{}:f,b=t.status;return r.a.createElement(K.a,{key:u,value:c,divider:!0},r.a.createElement(M.a,null,r.a.createElement(ce.a,{edge:"start",className:e.icon})),r.a.createElement(J.a,{container:!0,direction:"column"},r.a.createElement(J.a,{item:!0},r.a.createElement(k.a,{variant:"body2",display:"block",color:"secondary"},"At\xa0",c?function(e){var t=new Date(e),a=t.getTime()-60*t.getTimezoneOffset()*1e3,n=new Date(a).toISOString().match(/^(.*)T(.*)\..*/);return r.a.createElement(r.a.Fragment,null,n[1],"\xa0",n[2])}(c):r.a.createElement("i",null,"unknown time")," via\xa0",l)),r.a.createElement(J.a,{item:!0},r.a.createElement(k.a,{variant:"body2",display:"block",color:"secondary"},r.a.createElement(ae.a,{component:"span",fontWeight:800},ge(m),"\xa0"),!m&&r.a.createElement("i",null,"(",s,")\xa0"),"says:")),r.a.createElement(J.a,{item:!0},Object.entries(d).map((function(e,t){var a=Object(i.a)(e,2),n=a[0],c=a[1],o=c.issuerPetname,l=c.pursePetname,u=c.brandRegKey,s=c.extent;return r.a.createElement(k.a,{key:"offer".concat(n),variant:"body1"},0===t?"Pay":r.a.createElement(r.a.Fragment,null,"and\xa0pay"),"\xa0",r.a.createElement(ae.a,{component:"span",fontWeight:800},s,"\xa0",ge(o)),!o&&r.a.createElement(r.a.Fragment,null,"\xa0",r.a.createElement("i",null,"(",u,")"))," from\xa0",r.a.createElement(ae.a,{component:"span",fontWeight:800},ge(l)))})),Object.entries(g).map((function(e,t){var a=Object(i.a)(e,2),n=a[0],c=a[1],o=c.issuerPetname,l=c.pursePetname,u=c.brandRegKey,s=c.extent;return r.a.createElement(k.a,{key:"offer".concat(n),variant:"body1"},0===t?Object.keys(d).length>0?r.a.createElement(r.a.Fragment,null,"to\xa0receive"):"Receive":r.a.createElement(r.a.Fragment,null,"and\xa0receieve"),"\xa0",r.a.createElement(ae.a,{component:"span",fontWeight:800},s,"\xa0",ge(o)),!o&&r.a.createElement(r.a.Fragment,null,"\xa0",r.a.createElement("i",null,"(",u,")"))," into\xa0",r.a.createElement(ae.a,{component:"span",fontWeight:800},ge(l)))})))),r.a.createElement(ne.a,{className:e.buttons},"decline"===b&&r.a.createElement(de,{variant:"outlined",label:"Declined"}),"rejected"===b&&r.a.createElement(de,{variant:"outlined",label:"Rejected"}),"accept"===b&&r.a.createElement(fe,{variant:"outlined",label:"Accepted"}),"pending"===b&&r.a.createElement(pe,{size:"small","aria-label":"Cancel",onClick:function(){return function(e){n({type:"CANCEL_OFFER",payload:e})}(u)}},r.a.createElement(le.a,null)),"cancel"===b&&r.a.createElement(de,{variant:"outlined",label:"Cancelled"}),!b&&r.a.createElement(r.a.Fragment,null,r.a.createElement(me,{size:"small","aria-label":"Decline",onClick:function(){return function(e){n({type:"REJECT_OFFER",payload:e})}(u)}},r.a.createElement(le.a,null)),r.a.createElement(Ee,{size:"small","aria-label":"Accept",onClick:function(){return function(e){n({type:"CONFIRM_OFFER",payload:e})}(u)}},r.a.createElement(ue.a,null)))))}))):r.a.createElement(k.a,{color:"inherit"},"No transactions."))}var ve=Object(j.a)((function(e){return{paper:Object(T.a)({marginTop:e.spacing(3),marginBottom:e.spacing(3),padding:e.spacing(2)},e.breakpoints.up(600+2*e.spacing(3)),{marginTop:e.spacing(6),marginBottom:e.spacing(6),padding:e.spacing(3)}),title:Object(T.a)({padding:e.spacing(0,0,4)},e.breakpoints.down("xs"),{display:"none"}),aside:{borderRightColor:e.palette.divider,borderRightStyle:"solid",borderRightWidth:1}}}));function Oe(){var e=ve();return r.a.createElement(B.a,{className:e.paper},r.a.createElement(k.a,{component:"h1",variant:"h4",align:"center",className:e.title},"Agoric Simple Wallet"),r.a.createElement(J.a,{container:!0,spacing:2},r.a.createElement(J.a,{item:!0,xs:12,md:4,className:e.aside},r.a.createElement(Q,null)),r.a.createElement(J.a,{item:!0,xs:12,md:8},r.a.createElement(be,null))))}var ye=Object(j.a)((function(e){return{layout:Object(T.a)({width:"auto",marginLeft:e.spacing(2),marginRight:e.spacing(2)},e.breakpoints.up(1e3+2*e.spacing(2)),{width:1e3,marginLeft:"auto",marginRight:"auto"})}}));var he=Object(o.a)({palette:{primary:{main:"#AB2328"},secondary:{main:"#508AA8"}},typography:{useNextVariants:!0}});Object(c.render)(r.a.createElement((function(e){var t=e.children,a=Object(n.useReducer)(h,{active:!1,connected:!1,account:null,purses:null,inbox:null}),c=Object(i.a)(a,2),o=c[0],l=c[1],u=o.active;return Object(n.useEffect)((function(){function e(e){if(e){var t=e.type,a=e.data;"walletUpdatePurses"===t&&l(v(JSON.parse(a))),"walletUpdateInbox"===t&&l(O(JSON.parse(a)))}}u?function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};g()||f(e)}({onConnect:function(){l({type:"SERVER_CONNECTED"}),p({type:"walletGetPurses"}).then(e),p({type:"walletGetInbox"}).then(e)},onDisconnect:function(){l({type:"SERVER_DISCONNECTED"}),l({type:"DEACTIVATE_CONNECTION"}),l(v(null)),l(O(null))},onMessage:function(t){e(JSON.parse(t))}}):b()}),[u]),r.a.createElement(C.Provider,{value:{state:o,dispatch:l}},t)}),null,r.a.createElement(l.a,{theme:he},r.a.createElement((function(){var e=ye();return r.a.createElement(r.a.Fragment,null,r.a.createElement(R.a,null),r.a.createElement(_,null,r.a.createElement(U,null)),r.a.createElement("main",{className:e.layout},r.a.createElement(Oe,null)))}),null))),document.getElementById("root"))}},[[54,1,2]]]);
//# sourceMappingURL=main.49aad819.chunk.js.map