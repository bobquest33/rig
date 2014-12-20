Rig
===

Rig is a non-caching HTTP Reverse Proxy which implements a SAML Authorization step, optionally passing SAML attributes through to the back-end as headers. It's suitable for usage between a system such as Nginx and a given web-application, dashboard, etc.

You can use it to give a Non-SAML speaking HTTP application the ability to undestand authentication and authorization as HTTP headers which you can use in your authorization routines with a little modification.
