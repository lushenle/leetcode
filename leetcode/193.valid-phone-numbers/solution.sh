#!/usr/bin/env bash
# solution 1
grep -Eo '^(\([0-9]{3}\) |[0-9]{3}-)[0-9]{3}-[0-9]{4}$' file.txt

# solution 2
sed -n -r '/^(\([0-9]{3}\) |[0-9]{3}-)[0-9]{3}-[0-9]{4}$/p' file.txt

# solution 3
grep -P "^(\([0-9]{3}\)\s|[0-9]{3}-)[0-9]{3}-[0-9]{4}$" file.txt
