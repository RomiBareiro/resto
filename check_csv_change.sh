#!/bin/bash

CSV_FILE="template/csv_info.csv"

# save current timestamp
current_timestamp=$(stat -c %Y "$CSV_FILE")

# if exists, save previous file timestamp
previous_timestamp=$(cat previous_timestamp.txt 2>/dev/null)

# Compare timestamps
if [ -n "$previous_timestamp" ] && [ "$current_timestamp" != "$previous_timestamp" ]; then
    echo "El archivo CSV ha cambiado."
fi

# save new timestamp
echo "$current_timestamp" > previous_timestamp.txt

## TODO IMPROVEMENTS:

# We should check if there are 6 hr of difference from our current_timestamp and the last download timestamp (when we download the file)
# if yes:
    # check if current_timestamp file is different than previous, if yes download and replace file. If not, continue.
# if not:
    # use current file and continue

# To do this feat, we should use a data base to store updated request timestamp, current file timestamp, new file timestamp
# currently we are using a txt file without the necessary info

# Also this bash file must be called from setup(), and also there we need to do all the necessary settings for DB