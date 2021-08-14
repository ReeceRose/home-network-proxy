# Reporting Tool

## Setup

1. Extract reporting-tool to the desired install location.
2. Setup the reporting tool to run once an hour, at the start of the hour.

```bash
# Create a new cron job
crontab -e
# Add the following line
0 * * * * [INSTALL_LOCATION]/reporting-tool
```
