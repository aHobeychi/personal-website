# Building a Telegram Notification System for Badminton Court Availabilities

Finding available badminton courts can be frustrating—especially when slots get taken quickly. To solve this, I built an automated system that checks the website for open badminton court slots and sends instant notifications via Telegram.

This post walks through how I implemented this using **GitHub Actions**, **Playwright** for web scraping, and **Telegram** for push notifications.



## System Overview

The system consists of two GitHub Actions that run on a schedule:

1. **Update Chat IDs**  
   Queries the Telegram API for chat IDs of users who’ve messaged the bot and stores them in Firestore.

2. **Badminton Booking Notification**  
   Scrapes the City's booking site, filters results based on user preferences, and sends a Telegram message if any courts are available.



## Collecting Telegram Chat IDs

Before we can send messages, we need the chat IDs of users who’ve interacted with the Telegram bot.

- The GitHub Action calls:  
  `https://api.telegram.org/bot<token>/getUpdates`

- It extracts each user's `chat.id` and `username`.

- The list of chat IDs is saved in a Firestore collection named `chat_ids`.

This step ensures that every user who sends a message to the bot becomes eligible for notifications.



## Scraping the Booking Website

A Playwright script performs the following:

- Navigates to the badminton court booking site.
- Selects filters like **location**, **time range**, **days**, and **price**.
- Scrapes the results for available courts.

Playwright was selected because of its ability to handle modern JavaScript-heavy web pages.



## Filtering and Detection

After gathering raw availability data, we filter based on environment-configured search criteria:

- **Time range** (e.g., 6:00 PM – 9:00 PM)
- **Preferred days** (e.g., Friday, Saturday)
- **Specific locations** (e.g., Ahuntsic, Saint-Laurent)
- **Maximum price**

Only matching court availabilities are considered for notifications.



## Sending Telegram Notifications

If any courts match the filter criteria, the script sends messages to users via the Telegram Bot API:

Endpoint used:

```
https://api.telegram.org/bot/sendMessage
```

Each message includes:

- Date and time of availability
- Location
- A direct link to the booking site

Messages are sent in batch using the chat IDs stored in Firestore.



## Tech Stack

| Component         | Technology          |
|------------------|---------------------|
| Scraping         | Playwright          |
| Messaging        | Telegram Bot API    |
| Storage          | Firebase Firestore  |
| Automation       | GitHub Actions      |
| Config/Secrets   | GitHub Secrets      |



## Why This Matters

- **Time Saver**: No more manually checking the booking site.
- **Real-time Alerts**: Be the first to know when a court is free.
- **Extensible**: Can easily support other sports or locations.



## Source Code

You can find the source code [on GitHub](https://github.com/aHobeychi/Badminton-Booker).  
Follow the instructions in the `README.md` to deploy your own copy.



## Conclusion

This project demonstrates how you can build a simple yet powerful automation system using cloud-native tools. With minimal effort, you can turn a tedious manual task into a real-time personal assistant that delivers actionable information directly to your phone.