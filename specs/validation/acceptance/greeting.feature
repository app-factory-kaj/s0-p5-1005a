Feature: Hello greeting

  @story-1
  Rule: A caller who supplies a name receives a greeting addressed to it

    Scenario: Greeting a named caller
      Given the greeter service is available
      When Priya calls GET /hello with name "Priya"
      Then the response is a JSON greeting whose message addresses "Priya"

  @story-2
  Rule: A caller who omits the name still receives a usable greeting

    Scenario: Greeting with no name supplied
      Given the greeter service is available
      When Priya calls GET /hello with no name
      Then the response is a JSON greeting whose message is a generic greeting
