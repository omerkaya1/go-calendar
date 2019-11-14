# Main test suite for go-calendar
# The following features must be tested:
########################################
#  1) Event creation;
#  2) Event retrieval;
#  3) Event updating;
#  4) Event deletion.
########################################
Feature: Wholesome calendar
	In order to test the go-calendar application
	As a user that operates with the service through API
	The service should be able to do the following

	Scenario: Create a new event
        Given everything is ok
        When I make a send request to store an event:
			"""
			{
			    "user_name": "Morty",
			    "event_name": "Another 10th adventure",
			    "note": "oh geez!",
			    "start_time": "Fri Nov  13 15:00:00 MSK 2019",
			    "end_time": "Fri Nov  13 16:30:00 MSK 2019"
			}
			"""
        Then I receive an event ID

	Scenario: Get created event
		Given I have the event ID
		When I request this event by its ID
		Then the server returns it
		And it matches the the one we submitted

	Scenario: Update created event
		Given I have the event ID
		When I update the start time of the created event with "Fri Nov 13 14:00:00 MSK 2019" by its ID
		Then the server returns an ID of the updated event
		And both IDs should match

	Scenario: Delete created event
		Given I have the event ID
		When I request the deletion of the created event by its ID
		Then the server returns a success message
