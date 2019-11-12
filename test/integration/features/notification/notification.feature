# Main test suite for go-calendar
# The following features must be tested:
########################################
#  1) Notification sending;
#  2) Notification receiving.
########################################
Feature: Super notification service
	In order to test the notification service
	As a consumer of notification messages
	I need to be able to see notification messages

	Scenario: Receive a notification
		Given everything is set up
		When an event is stored in the DB
		And the notification service is started
		Then the notification service returns stored event as a message
