# Hypertext Transfer Protocol (HTTP/1.1): Semantics and Content

## Section 6.4.2 - 301 Moved Permanently

The 301 (Moved Permanently) status code indicates that the target
resource has been assigned a new permanent URI and any future
references to this resource ought to use one of the enclosed URIs.
Clients with link-editing capabilities ought to automatically re-link
references to the effective request URI to one or more of the new
references sent by the server, where possible.

The server SHOULD generate a Location header field in the response
containing a preferred URI reference for the new permanent URI.  The
user agent MAY use the Location field value for automatic
redirection.  The server's response payload usually contains a short
hypertext note with a hyperlink to the new URI(s).

Note: For historical reasons, a user agent MAY change the request
method from POST to GET for the subsequent request.  If this
behavior is undesired, the 307 (Temporary Redirect) status code
can be used instead.

A 301 response is cacheable by default; i.e., unless otherwise
indicated by the method definition or explicit cache controls (see
Section 4.2.2 of [RFC7234](https://datatracker.ietf.org/doc/html/rfc7234#section-4.2.2)).

[Reference](https://datatracker.ietf.org/doc/html/rfc7231#section-6.4.2)
