# Comms API Challenge

The challenge is to solve an emerging business problem within an existing codebase.

**Context:**

The existing codebase represents a restful API for various comms operations. The API utilizes an email service inside the route handlers to send emails. In the `main.go` server application the email service uses a sendgrid implementation, but the handlers also have unit tests which utilize a mock implementation to ensure full coverage of the handler email functionality.

**Business Problem:**

The comms api needs to support a new operation for adding policy coverage. The new `add-policy-coverage` route should have the same data contract as the others, with the exception that the payload for the new route also needs to support `CC` in addition to `To` for the email recipients. Until this point only `To` recipients have been needed, so the current email service provider does not yet support `CC`.

**Instructions:** 

* [ ] Create a handler for the new comms operation and implement the correct business logic inside of the handler.
* [ ] Make whatever improvements are needed to support the new email `CC` field.
* [ ] Create a unit test for the new handler to prove it meets business requirements.
