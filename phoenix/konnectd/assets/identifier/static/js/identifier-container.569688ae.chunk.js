(window.webpackJsonp=window.webpackJsonp||[]).push([[1],{440:function(e,t,n){e.exports=n.p+"static/media/kopano-logo.10e256c7.svg"},551:function(e,t,n){"use strict";n.r(t);var a=n(32),r=n(35),o=n(38),c=n(33),i=n(36),s=n(8),l=n.n(s),u=n(37),m=n(86),d=n(87),p=n(25);function h(e){return e.state=Math.random().toString(36).substring(7),e}function f(e,t){var n,a,r=arguments.length>2&&void 0!==arguments[2]&&arguments[2];if(Array.prototype.some.call(document.styleSheets,function(r){return Array.prototype.some.call(r.cssRules,function(o){return a=r.href,e===o.selectorText&&Array.prototype.some.call(o.style,function(e){return t===e&&(n=o.style.getPropertyValue(t),!0)})})}),n&&r){if(!(n=n.match(/(?:\(['|"]?)(.*?)(?:['|"]?\))/)[1]))return null;if(a){var o=a.replace(/\\/g,"/").replace(/\/[^/]*$/,"");n=o+"/"+n}}return n}var g=n(185),b=n(68),v=n(12),y=n.n(v),E=n(206),w=n.n(E),k=n(477),j=n.n(k),O=n(66),N=n.n(O),P=n(480),M=n(406),S=n.n(M),_=n(426);var C=n.n(_)()()(function(e){return l.a.createElement(S.a,e)}),x=n(440),B=n.n(x),A=n(29),D=n(461),F=n.n(D),H=n(46),L=n.n(H),T=n(55),R=n.n(T),K=n(83),I=n.n(K),U=n(200),q=n.n(U),W=n(57);function G(e,t){var n={};t.prompt&&(n.prompt=t.prompt);var a=e;switch(e){case"oauth":case"consent":case"oidc":n.scope=t.scope||"",n.client_id=t.client_id||"",n.redirect_uri=t.redirect_uri||"",t.id_token_hint&&(n.id_token_hint=t.id_token_hint),t.max_age&&(n.max_age=t.max_age),t.claims_scope&&(n.scope+=" "+t.claims_scope);break;default:a=null}return a&&(n.flow=a),n}var z,X=n(186),J=n(547),Y=n(70),V="konnect.error.login.validate.missingUsername",$="konnect.error.login.validate.missingPassword",Q="konnect.error.login.failed",Z="konnet.error.http.networkError",ee="konnect.error.http.unexpectedResponseStatus",te="konnect.error.http.unexpectedResponseState",ne=Object(A.d)((z={},Object(Y.a)(z,V,{id:V,defaultMessage:"Enter an username"}),Object(Y.a)(z,$,{id:$,defaultMessage:"Enter a password"}),Object(Y.a)(z,Q,{id:Q,defaultMessage:"Logon failed. Please verify your credentials and try again."}),Object(Y.a)(z,Z,{id:Z,defaultMessage:"Network error. Please check your connection and try again."}),Object(Y.a)(z,ee,{id:ee,defaultMessage:"Unexpected HTTP response: {status}. Please check your connection and try again."}),Object(Y.a)(z,te,{id:te,defaultMessage:"Unexpected response state: {state}"}),z)),ae=function(e){function t(e,n){var r;return Object(a.a)(this,t),(r=Object(o.a)(this,Object(c.a)(t).call(this,e))).values=void 0,void 0!==Error.captureStackTrace&&Error.captureStackTrace(Object(X.a)(Object(X.a)(r)),t),r.values=n,r}return Object(i.a)(t,e),t}(Object(J.a)(Error));var re=Object(A.e)(function(e){var t=e.error,n=e.intl;if(!t)return null;var a=t.id?t.id:t.message,r=Object.assign({},{id:a,defaultMessage:t.id?t.message:void 0},ne[a]);return n.formatMessage(r,t.values)});function oe(e){return e.request&&(e=e.response?new ae(ee,e.response):new ae(Z)),e}function ce(e){return{type:W.b,error:e}}function ie(){return{type:W.k}}function se(e){var t=e.success,n=e.username,a=e.displayName;return{type:W.c,state:!0===t,username:n,displayName:a,hello:e}}function le(){return function(e,t){e(ie());var n=t().common,a=h(G(n.flow,n.query));return q.a.post("./identifier/_/hello",a,{headers:{"Kopano-Konnect-XSRF":"1"}}).then(function(e){switch(e.status){case 200:return e.data;case 204:return{success:!1,state:e.headers["kopano-konnect-state"]};default:throw new ae(ee,e)}}).then(function(t){if(t.state!==a.state)throw new ae(te,t);return e(se(t)),Promise.resolve(t)}).catch(function(t){t=oe(t),e(ce(t))})}}function ue(){return function(e){e(ie()),e({type:W.i});var t=h({});return q.a.post("./identifier/_/logoff",t,{headers:{"Kopano-Konnect-XSRF":"1"}}).then(function(e){switch(e.status){case 200:return e.data;default:throw new ae(ee,e)}}).then(function(n){if(n.state!==t.state)throw new ae(te,n);var a;return e((a=!0===n.success,{type:W.d,state:a})),Promise.resolve(n)}).catch(function(t){t=oe(t),e(ce(t))})}}var me=function(e){function t(){return Object(a.a)(this,t),Object(o.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(i.a)(t,e),Object(r.a)(t,[{key:"render",value:function(){var e=this,t=this.props,n=t.classes,a=t.error;return l.a.createElement(w.a,{container:!0,direction:"column",alignItems:"center",justify:"center",spacing:0,className:n.root},l.a.createElement(w.a,{item:!0,align:"center"},I()(null===a)(function(){return l.a.createElement(F.a,{className:n.progress})}),I()(null!==a)(function(){return l.a.createElement("div",null,l.a.createElement(L.a,{variant:"h5",gutterBottom:!0,align:"center"},l.a.createElement(A.a,{id:"konnect.loading.error.headline",defaultMessage:"Failed to connect to server"})),l.a.createElement(L.a,{gutterBottom:!0,align:"center",color:"error"},l.a.createElement(re,{error:a})),l.a.createElement(R.a,{autoFocus:!0,color:"primary",variant:"outlined",className:n.button,onClick:function(t){return e.retry(t)}},l.a.createElement(A.a,{id:"konnect.login.retryButton.label",defaultMessage:"Retry"})))})))}},{key:"retry",value:function(e){e.preventDefault(),this.props.dispatch(function(e){return e(ce(null)),e(le())})}}]),t}(s.Component),de=Object(u.b)(function(e){return{error:e.common.error}})(Object(p.withStyles)(function(e){return{root:{flexGrow:1,position:"absolute",top:0,bottom:0,left:0,right:0},progress:{height:"4px",width:"100px"},button:{marginTop:5*e.spacing.unit}}})(me));function pe(e){var t=e.classes,n=e.withoutLogo,a=e.withoutPadding,r=e.loading,o=e.children,c=e.className,i=e.DialogProps,s=e.PaperProps,u=Object(b.a)(e,["classes","withoutLogo","withoutPadding","loading","children","className","DialogProps","PaperProps"]),m=n?null:l.a.createElement(j.a,null,l.a.createElement("img",{src:B.a,className:t.logo,alt:"Kopano"})),d=r?l.a.createElement(de,null):a?o:l.a.createElement(j.a,null,o);return l.a.createElement(w.a,Object.assign({container:!0,justify:"center",alignItems:"center",spacing:0,className:y()(t.root,c)},u),l.a.createElement(C,Object.assign({open:!0,fullWidth:!0,maxWidth:"sm",disableBackdropClick:!0,disableEscapeKeyDown:!0,hideBackdrop:!0},i,{PaperProps:Object(g.a)({elevation:4},s)}),l.a.createElement("div",{className:t.content},m,d),l.a.createElement(N.a,{className:t.actions,disableActionSpacing:!0},l.a.createElement(P.a,{disableUnderline:!0}))))}pe.defaultProps={withoutLogo:!1,withoutPadding:!1,loading:!1};var he=Object(p.withStyles)(function(e){return{root:{display:"flex",flex:1},content:{paddingTop:24,paddingBottom:12,minHeight:500,paddingLeft:2*e.spacing.unit,paddingRight:2*e.spacing.unit,position:"relative"},logo:{height:18,marginBottom:2*e.spacing.unit},actions:{marginTop:-40,justifyContent:"flex-start",paddingLeft:3*e.spacing.unit,paddingRight:3*e.spacing.unit}}})(pe),fe=n(215),ge=n.n(fe),be=n(214),ve=n.n(be),ye=n(505),Ee=n.n(ye),we=n(135),ke=n.n(we),je="0",Oe="1";function Ne(e){return{type:W.f,errors:e}}function Pe(e,t){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:Oe;return function(a,r){a(function(e,t){return{type:W.j,username:e,password:t}}(e,t)),a(se({username:e}));var o=r().common,c=o.flow,i=o.query,s=[];switch(n){case Oe:s.push(e,t,n);break;case je:s.push(e,"",n)}var l=h({params:s,hello:G(c,i)});return q.a.post("./identifier/_/logon",l,{headers:{"Kopano-Konnect-XSRF":"1"}}).then(function(e){switch(e.status){case 200:return e.data;case 204:return{success:!1,state:e.headers["kopano-konnect-state"],errors:{http:new Error(Q)}};default:throw new ae(ee,e)}}).then(function(t){if(t.state!==l.state)throw new ae(te,t);var n=t.hello;return n||(n={success:t.success,username:e}),a(se(n)),a(function(e){var t=e.success,n=e.errors;return{type:W.e,success:t,errors:n}}(t)),Promise.resolve(t)}).catch(function(e){var t={http:e=oe(e)};return a(Ne(t)),{success:!1,errors:t}})}}function Me(){var e=arguments.length>0&&void 0!==arguments[0]&&arguments[0],t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return function(n,a){n(function(){return{type:arguments.length>0&&void 0!==arguments[0]&&arguments[0]?W.g:W.h}}(e));var r=a().common.query,o=h({allow:e,scope:t,client_id:r.client_id||"",redirect_uri:r.redirect_uri||"",ref:r.state||"",flow_nonce:r.nonce||""});return q.a.post("./identifier/_/consent",o,{headers:{"Kopano-Konnect-XSRF":"1"}}).then(function(e){switch(e.status){case 200:return e.data;case 204:return{success:!0,state:e.headers["kopano-konnect-state"]};default:throw new ae(ee,e)}}).then(function(e){if(e.state!==o.state)throw new ae(te,e);return n(function(e){var t=e.success,n=e.errors;return{type:W.a,success:t,errors:n}}(e)),Promise.resolve(e)}).catch(function(e){var t={http:e=oe(e)};return n(Ne(t)),{success:!1,errors:t}})}}function Se(e,t,n){return function(a){return a(function(e,t,n){return function(a){return new Promise(function(r,o){var c={};e||(c.username=new Error(V)),t||n||(c.password=new Error($)),a(Ne(c)),0===Object.keys(c).length?r(c):o(c)})}}(e,t,n)).then(function(){return a(Pe(e,t,n?je:Oe))}).catch(function(e){return{success:!1,errors:e}})}}function _e(e,t){var n=arguments.length>2&&void 0!==arguments[2]&&arguments[2],a=arguments.length>3&&void 0!==arguments[3]?arguments[3]:{};return function(r,o){if(e){var c=o().common,i=c.flow,s=c.query,l=c.hello,u=Object.assign({},s,a);switch(i){case"oauth":case"consent":case"oidc":if(l.details.flow!==i)break;if(!n&&"consent"===l.details.next)return void t.replace("/consent".concat(t.location.search).concat(t.location.hash));if(l.details.continue_uri)return u.prompt="none",void window.location.replace(l.details.continue_uri+"?"+ke.a.stringify(u));break;default:if(u.continue&&0===u.continue.indexOf(document.location.origin))return void window.location.replace(u.continue)}var m="/welcome";"REPLACE"===t.action&&(m=m+t.location.search+t.location.hash),r(Ne({})),t.push(m)}}}var Ce=function(e){function t(){var e,n;Object(a.a)(this,t);for(var r=arguments.length,i=new Array(r),s=0;s<r;s++)i[s]=arguments[s];return(n=Object(o.a)(this,(e=Object(c.a)(t)).call.apply(e,[this].concat(i)))).state={},n.onAutoFill=function(e,t){n.setState(Object(Y.a)({},"autoFill-".concat(e),!!t||void 0))},n.bindAutoFill=function(e){return function(t){t&&t.addEventListener&&t.addEventListener("animationstart",function(t){switch(t.animationName){case"onAutoFillStart":return n.onAutoFill(e,!0);case"onAutoFillCancel":return n.onAutoFill(e,!1)}})}},n}return Object(i.a)(t,e),Object(r.a)(t,[{key:"componentDidMount",value:function(){var e=this.props,t=e.hello,n=e.query,a=e.dispatch,r=e.history;if(t&&t.state&&"PUSH"!==r.action)return"select_account"!==n.prompt?void a(_e(!0,r)):void r.replace("/chooseaccount".concat(r.location.search).concat(r.location.hash))}},{key:"render",value:function(){var e=this,t=this.props,n=t.loading,a=t.errors,r=t.classes,o=t.username,c={username:{autoCapitalize:"off",spellCheck:"false",className:r.input,ref:this.bindAutoFill("username")},password:{className:r.input,ref:this.bindAutoFill("password")}};return l.a.createElement("div",null,l.a.createElement(L.a,{variant:"h5",component:"h3"},l.a.createElement(A.a,{id:"konnect.login.headline",defaultMessage:"Sign in"})),l.a.createElement(L.a,{variant:"subtitle1",className:r.subHeader},l.a.createElement(A.a,{id:"konnect.login.subHeader",defaultMessage:"with your Kopano account"})),l.a.createElement("form",{action:"",onSubmit:function(t){return e.logon(t)}},l.a.createElement("div",null,l.a.createElement(Ee.a,{label:l.a.createElement(A.a,{id:"konnect.login.usernameField.label",defaultMessage:"Username"}),error:!!a.username,helperText:l.a.createElement(re,{error:a.username}),fullWidth:!0,margin:"dense",variant:"outlined",InputLabelProps:{shrink:this.state["autoFill-username"]},autoFocus:!0,inputProps:c.username,value:o,onChange:this.handleChange("username"),autoComplete:"kopano-account username"}),l.a.createElement(Ee.a,{type:"password",label:l.a.createElement(A.a,{id:"konnect.login.passwordField.label",defaultMessage:"Password"}),error:!!a.password,helperText:l.a.createElement(re,{error:a.password}),fullWidth:!0,margin:"dense",InputLabelProps:{shrink:this.state["autoFill-password"]},inputProps:c.password,variant:"outlined",onChange:this.handleChange("password"),autoComplete:"kopano-account current-password"}),l.a.createElement(N.a,null,l.a.createElement("div",{className:r.wrapper},l.a.createElement(R.a,{type:"submit",color:"primary",variant:"contained",className:r.button,disabled:!!n,onClick:function(t){return e.logon(t)}},l.a.createElement(A.a,{id:"konnect.login.nextButton.label",defaultMessage:"Next"})),n&&l.a.createElement(ge.a,{size:24,className:r.buttonProgress})))),I()(a.http)(function(){return l.a.createElement(L.a,{variant:"subtitle2",color:"error",className:r.message},l.a.createElement(re,{error:a.http}))})))}},{key:"handleChange",value:function(e){var t=this;return function(n){t.props.dispatch(function(e,t){return{type:W.m,name:e,value:t}}(e,n.target.value))}}},{key:"logon",value:function(e){e.preventDefault();var t=this.props,n=t.username,a=t.password,r=t.dispatch,o=t.history;r(Se(n,a,!1)).then(function(e){e.success&&r(_e(e.success,o))})}}]),t}(s.Component),xe=Object(u.b)(function(e){var t=e.login,n=t.loading,a=t.username,r=t.password,o=t.errors,c=e.common;return{loading:n,username:a,password:r,errors:o,hello:c.hello,query:c.query}})(Object(p.withStyles)(function(e){return{button:{margin:e.spacing.unit,minWidth:100},buttonProgress:{color:ve.a[500],position:"absolute",top:"50%",left:"50%",marginTop:-12,marginLeft:-12},subHeader:{marginBottom:3*e.spacing.unit},wrapper:{position:"relative",display:"inline-block"},message:{marginTop:2*e.spacing.unit,marginBottom:2*e.spacing.unit},input:{"&:-webkit-autofill":{animationName:"onAutoFillStart"},"&:not(:-webkit-autofill)":{animationName:"onAutoFillCancel"}}}})(Ce)),Be=n(122),Ae=n.n(Be),De=n(128),Fe=n.n(De),He=n(216),Le=n.n(He),Te=n(521),Re=n.n(Te),Ke=function(e){function t(){return Object(a.a)(this,t),Object(o.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(i.a)(t,e),Object(r.a)(t,[{key:"componentDidMount",value:function(){var e=this.props,t=e.hello,n=e.history;t&&t.state||"PUSH"===n.action||n.replace("/identifier".concat(n.location.search).concat(n.location.hash))}},{key:"render",value:function(){var e=this,t=this.props,n=t.loading,a=t.errors,r=t.classes,o=t.hello,c=null;a.http&&(c=l.a.createElement(L.a,{color:"error",className:r.message},l.a.createElement(re,{error:a.http})));var i="";return o&&o.state&&(i=o.username),l.a.createElement("div",null,l.a.createElement(L.a,{variant:"h5",component:"h3"},l.a.createElement(A.a,{id:"konnect.chooseaccount.headline",defaultMessage:"Choose an account"})),l.a.createElement(L.a,{variant:"subtitle1",className:r.subHeader},l.a.createElement(A.a,{id:"konnect.chooseaccount.subHeader",defaultMessage:"to sign in to Kopano"})),l.a.createElement("form",{action:"",onSubmit:function(t){return e.logon(t)}},l.a.createElement(Ae.a,{disablePadding:!0,className:r.accountList},l.a.createElement(Fe.a,{button:!0,disableGutters:!0,className:r.accountListItem,disabled:!!n,onClick:function(t){return e.logon(t)}},l.a.createElement(Re.a,null,i.substr(0,1)),l.a.createElement(Le.a,{primary:i})),l.a.createElement(Fe.a,{button:!0,disableGutters:!0,className:r.accountListItem,disabled:!!n,onClick:function(t){return e.logoff(t)}},l.a.createElement(Re.a,null,l.a.createElement(A.a,{id:"konnect.chooseaccount.useOther.persona.label",defaultMessage:"?"})),l.a.createElement(Le.a,{primary:l.a.createElement(A.a,{id:"konnect.chooseaccount.useOther.label",defaultMessage:"Use another account"})}))),c))}},{key:"logon",value:function(e){e.preventDefault();var t=this.props,n=t.hello,a=t.dispatch,r=t.history;a(Se(n.username,"",!0)).then(function(e){e.success&&a(_e(e.success,r))})}},{key:"logoff",value:function(e){e.preventDefault();var t=this.props.history;t.push("/identifier".concat(t.location.search).concat(t.location.hash))}}]),t}(s.Component),Ie=Object(u.b)(function(e){var t=e.login;return{loading:t.loading,errors:t.errors,hello:e.common.hello}})(Object(p.withStyles)(function(e){return{subHeader:{marginBottom:2*e.spacing.unit},message:{marginTop:2*e.spacing.unit},accountList:{marginLeft:-3*e.spacing.unit,marginRight:-3*e.spacing.unit},accountListItem:{paddingLeft:3*e.spacing.unit,paddingRight:3*e.spacing.unit}}})(Ke)),Ue=n(532),qe=n.n(Ue),We=function(e){var t=e.client,n=Object(b.a)(e,["client"]);return l.a.createElement("span",n,t.display_name?t.display_name:t.id)},Ge=n(524),ze=n.n(Ge),Xe=Object(A.d)({scope_alias_basic:{id:"konnect.scopeDescription.aliasBasic",defaultMessage:"Access your basic account information"},scope_offline_access:{id:"konnect.scopeDescription.offlineAccess",defaultMessage:"Keep the allowed access persistently and forever"}}),Je=Object(p.withStyles)(function(){return{row:{paddingTop:0,paddingBottom:0}}})(Object(A.e)(function(e){var t=e.scopes,n=e.meta,a=e.classes,r=e.intl,o=Object(b.a)(e,["scopes","meta","classes","intl"]),c=n.mapping,i=n.definitions,s=[],u={};for(var m in t)if(t[m]){var d=c[m];if(d){if(u[d])continue;u[d]=!0}else d=m;var p=i[d],h=void 0;if(p){if(p.id){var f=Xe[p.id];f&&(h=r.formatMessage(f))}h||(h=p.description)}h||(h=l.a.createElement(A.a,{id:"konnect.scopeDescription.scope",defaultMessage:"Scope: {scope}",values:{scope:m}})),s.push(l.a.createElement(Fe.a,{disableGutters:!0,dense:!0,key:d,className:a.row},l.a.createElement(ze.a,{checked:!0,disableRipple:!0,disabled:!0}),l.a.createElement(Le.a,{primary:h})))}return l.a.createElement(Ae.a,o,s)})),Ye=function(e){function t(){var e,n;Object(a.a)(this,t);for(var r=arguments.length,i=new Array(r),s=0;s<r;s++)i[s]=arguments[s];return(n=Object(o.a)(this,(e=Object(c.a)(t)).call.apply(e,[this].concat(i)))).action=function(){var e=arguments.length>0&&void 0!==arguments[0]&&arguments[0],t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return function(a){if(a.preventDefault(),void 0!==e){var r=Object.keys(t).filter(function(e){return!!t[e]}).join(" "),o=n.props,c=o.dispatch,i=o.history;c(Me(e,r)).then(function(e){e.success&&c(_e(e.success,i,!0,{konnect:e.state}))})}}},n}return Object(i.a)(t,e),Object(r.a)(t,[{key:"componentDidMount",value:function(){var e=this.props,t=e.dispatch,n=e.hello,a=e.history,r=e.client;n&&n.state&&r||"PUSH"===a.action||a.replace("/identifier".concat(a.location.search).concat(a.location.hash)),t(Ne({}))}},{key:"render",value:function(){var e=this.props,t=e.classes,n=e.loading,a=e.hello,r=e.errors,o=e.client,c=a.details.scopes||{},i=a.details.meta||{};return l.a.createElement("div",null,l.a.createElement(L.a,{variant:"h5",component:"h3"},l.a.createElement(A.a,{id:"konnect.consent.headline",defaultMessage:"Hi {displayName}",values:{displayName:a.displayName}})),l.a.createElement(L.a,{variant:"subtitle1",className:t.subHeader},a.username),l.a.createElement(L.a,{variant:"subtitle1",gutterBottom:!0},l.a.createElement(A.a,{id:"konnect.consent.message",defaultMessage:"{clientDisplayName} wants to",values:{clientDisplayName:l.a.createElement(qe.a,{placement:"bottom",title:l.a.createElement(A.a,{id:"konnect.consent.tooltip.client",defaultMessage:'Clicking "Allow" will redirect you to: {redirectURI}',values:{redirectURI:o.redirect_uri}})},l.a.createElement("em",null,l.a.createElement(We,{client:o})))}})),l.a.createElement(Je,{dense:!0,disablePadding:!0,className:t.scopesList,scopes:c,meta:i.scopes}),l.a.createElement(L.a,{variant:"subtitle1",gutterBottom:!0},l.a.createElement(A.a,{id:"konnect.consent.question",defaultMessage:"Allow {clientDisplayName} to do this?",values:{clientDisplayName:l.a.createElement("em",null,l.a.createElement(We,{client:o}))}})),l.a.createElement(L.a,{color:"secondary"},l.a.createElement(A.a,{id:"konnect.consent.consequence",defaultMessage:"By clicking Allow, you allow this app to use your information."})),l.a.createElement("form",{action:"",onSubmit:this.action(void 0,c)},l.a.createElement(N.a,null,l.a.createElement("div",{className:t.wrapper},l.a.createElement(R.a,{color:"secondary",className:t.button,disabled:!!n,onClick:this.action(!1,c)},l.a.createElement(A.a,{id:"konnect.consent.cancelButton.label",defaultMessage:"Cancel"})),n&&n!==W.g&&l.a.createElement(ge.a,{size:24,className:t.buttonProgress})),l.a.createElement("div",{className:t.wrapper},l.a.createElement(R.a,{type:"submit",color:"primary",variant:"contained",className:t.button,disabled:!!n,onClick:this.action(!0,c)},l.a.createElement(A.a,{id:"konnect.consent.allowButton.label",defaultMessage:"Allow"})),n===W.g&&l.a.createElement(ge.a,{size:24,className:t.buttonProgress}))),I()(r.http)(function(){return l.a.createElement(L.a,{variant:"subtitle2",color:"error",className:t.message},l.a.createElement(re,{error:r.http}))})))}}]),t}(s.Component),Ve=Object(u.b)(function(e){var t=e.common.hello,n=e.login;return{loading:n.loading,errors:n.errors,hello:t,client:t.details.client||{}}})(Object(p.withStyles)(function(e){return{button:{margin:e.spacing.unit,minWidth:100},buttonProgress:{color:ve.a[500],position:"absolute",top:"50%",left:"50%",marginTop:-12,marginLeft:-12},subHeader:{marginBottom:2*e.spacing.unit},scopesList:{marginBottom:2*e.spacing.unit},wrapper:{marginTop:2*e.spacing.unit,position:"relative",display:"inline-block"},message:{marginTop:2*e.spacing.unit,marginBottom:2*e.spacing.unit}}})(Ye)),$e=Object(d.f)(function(e){var t=e.target,n=e.location,a=Object(b.a)(e,["target","location"]),r={pathname:t,search:n.search,hash:n.hash};return l.a.createElement(d.a,Object.assign({to:r},a))}),Qe=function(e){function t(){return Object(a.a)(this,t),Object(o.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(i.a)(t,e),Object(r.a)(t,[{key:"componentDidMount",value:function(){this.props.dispatch(le())}},{key:"render",value:function(){var e=null===this.props.hello;return l.a.createElement(he,{loading:e},l.a.createElement(d.d,null,l.a.createElement(d.b,{path:"/identifier",exact:!0,component:xe}),l.a.createElement(d.b,{path:"/chooseaccount",exact:!0,component:Ie}),l.a.createElement(d.b,{path:"/consent",exact:!0,component:Ve}),l.a.createElement($e,{target:"/identifier"})))}}]),t}(l.a.PureComponent),Ze=Object(u.b)(function(e){return{hello:e.common.hello}})(Object(p.withStyles)(function(){return{}})(Qe)),et=function(e){function t(){return Object(a.a)(this,t),Object(o.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(i.a)(t,e),Object(r.a)(t,[{key:"render",value:function(){var e=this,t=this.props,n=t.classes,a=t.hello,r=null===a;return l.a.createElement(he,{loading:r},l.a.createElement(L.a,{variant:"h5",component:"h3"},l.a.createElement(A.a,{id:"konnect.welcome.headline",defaultMessage:"Welcome {displayName}",values:{displayName:a.displayName}})),l.a.createElement(L.a,{variant:"subtitle1",className:n.subHeader},a.username),l.a.createElement(L.a,{gutterBottom:!0},l.a.createElement(A.a,{id:"konnect.welcome.message",defaultMessage:"You are signed in - awesome!"})),l.a.createElement(N.a,null,l.a.createElement(R.a,{color:"secondary",className:n.button,variant:"contained",onClick:function(t){return e.logoff(t)}},l.a.createElement(A.a,{id:"konnect.welcome.signoutButton.label",defaultMessage:"Sign out"}))))}},{key:"logoff",value:function(e){var t=this;e.preventDefault(),this.props.dispatch(ue()).then(function(e){var n=t.props.history;e.success&&n.push("/identifier")})}}]),t}(l.a.PureComponent),tt=Object(u.b)(function(e){return{hello:e.common.hello}})(Object(p.withStyles)(function(e){return{button:{margin:e.spacing.unit,minWidth:100},subHeader:{marginBottom:5*e.spacing.unit}}})(et)),nt=function(e){function t(){return Object(a.a)(this,t),Object(o.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(i.a)(t,e),Object(r.a)(t,[{key:"componentDidMount",value:function(){this.props.dispatch(le())}},{key:"render",value:function(){var e=this,t=this.props,n=t.classes,a=t.hello,r=null===a;return l.a.createElement(he,{loading:r},I()(null!==a&&!a.state)(function(){return l.a.createElement("div",null,l.a.createElement(L.a,{variant:"h5",component:"h3"},l.a.createElement(A.a,{id:"konnect.goodbye.headline",defaultMessage:"Goodbye"})),l.a.createElement(L.a,{variant:"subtitle1",className:n.subHeader},l.a.createElement(A.a,{id:"konnect.goodbye.subHeader",defaultMessage:"you have been signed out from your Kopano account"})),l.a.createElement(L.a,{gutterBottom:!0},l.a.createElement(A.a,{id:"konnect.goodbye.message.close",defaultMessage:"You can close this window now."})))}),I()(null!==a&&!0===a.state)(function(){return l.a.createElement("div",null,l.a.createElement(L.a,{variant:"h5",component:"h3"},l.a.createElement(A.a,{id:"konnect.goodbye.confirm.headline",defaultMessage:"Hello {displayName}",values:{displayName:a.displayName}})),l.a.createElement(L.a,{variant:"subtitle1",className:n.subHeader},l.a.createElement(A.a,{id:"konnect.goodbye.confirm.subHeader",defaultMessage:"please confirm sign out"})),l.a.createElement(L.a,{gutterBottom:!0},l.a.createElement(A.a,{id:"konnect.goodbye.message.confirm",defaultMessage:"Press the button below, to sign out from your Kopano account now."})),l.a.createElement(N.a,null,l.a.createElement("div",{className:n.wrapper},l.a.createElement(R.a,{color:"secondary",className:n.button,onClick:function(t){return e.logoff(t)}},l.a.createElement(A.a,{id:"konnect.goodbye.signoutButton.label",defaultMessage:"Sign out"})))))}))}},{key:"logoff",value:function(e){var t=this;e.preventDefault(),this.props.dispatch(ue()).then(function(e){var n=t.props.history;e.success&&(t.props.dispatch(le()),n.push("/goodbye"))})}}]),t}(l.a.PureComponent),at=Object(u.b)(function(e){return{hello:e.common.hello}})(Object(p.withStyles)(function(e){return{subHeader:{marginBottom:5*e.spacing.unit},wrapper:{marginTop:5*e.spacing.unit,position:"relative",display:"inline-block"}}})(nt)),rt=function(e){var t=e.component,n=e.hello,a=Object(b.a)(e,["component","hello"]);return l.a.createElement(d.b,Object.assign({},a,{render:function(e){return n?l.a.createElement(t,e):l.a.createElement($e,{target:"/identifier"})}}))};!function(){var e=f("#bg-enhanced.enhanced","background-image",!0),t=f("#bg-enhanced.enhanced::after","background-image",!0),n=[];e&&n.push(new Promise(function(t){var n=new Image;n.onload=function(){t()},n.src=e})),t&&n.push(new Promise(function(e){var n=new Image;n.onload=function(){e()},n.src=t})),Promise.all(n).then(function(){window.document.getElementById("bg-enhanced").className+=" enhanced"})}();var ot=function(e){function t(){return Object(a.a)(this,t),Object(o.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(i.a)(t,e),Object(r.a)(t,[{key:"render",value:function(){var e=this.props,t=e.classes,n=e.hello,a=e.pathPrefix;return l.a.createElement("div",{className:t.root},l.a.createElement(m.a,{basename:a},l.a.createElement(d.d,null,l.a.createElement(rt,{path:"/welcome",exact:!0,component:tt,hello:n}),l.a.createElement(d.b,{path:"/goodbye",exact:!0,component:at}),l.a.createElement(d.b,{path:"/",component:Ze}))))}},{key:"reload",value:function(e){e.preventDefault(),window.location.reload()}}]),t}(s.PureComponent);t.default=Object(u.b)(function(e){var t=e.common;return{hello:t.hello,updateAvailable:t.updateAvailable,pathPrefix:t.pathPrefix}})(Object(p.withStyles)(function(){return{root:{position:"relative",display:"flex",flex:1}}})(ot))}}]);
//# sourceMappingURL=identifier-container.569688ae.chunk.js.map