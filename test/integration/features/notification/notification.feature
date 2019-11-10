# Main test suite for go-calendar
# The following features must be tested:
########################################
#  1) Notification sending;
#  2) Notification receiving.
########################################
Feature: Get notification about the event
	In order to test the notification service
	As a consumer of notification messages
	I need to be able to see notification messages

	Scenario: Receive a notification
		Given a new event is stored in the DB:
		"""
			{
			    "user_name": "morty",
			    "event_name": "Another 10th adventure",
			    "note": "oh geez! Vindicators 4, bitch!",
			    "start_time": "Fri Nov  11 15:00:00 MSK 2019",
			    "end_time": "Fri Nov  11 16:30:00 MSK 2019"
			}
		"""
		And the notification service is started
		Then the notification service returns created event as a message
