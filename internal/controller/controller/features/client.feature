Feature: Client management
    In order to make an order
    As a client
    I need to be able to identify using my document

    Scenario: Get client details
        When I send a GET request to "/clients"
        Then Status code should be 200
        And Client details should be returned