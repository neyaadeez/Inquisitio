import re
from datetime import datetime, timedelta
from collections import defaultdict

log_pattern = re.compile(
    r"(?P<timestamp>\w+\s+\d+\s+\d+:\d+:\d+) .* Failed password for .* from (?P<ip>\d+\.\d+\.\d+\.\d+)"
)

def parse_timestamp(timestamp_str):
    return datetime.strptime(timestamp_str, "%b %d %H:%M:%S")

log_file_path = "auth.log"  
attempt_threshold = 5  
time_window = timedelta(minutes=5)  

failed_attempts = defaultdict(list)

with open(log_file_path, "r") as log_file:
    for line in log_file:
        match = log_pattern.search(line)
        if match:
            timestamp_str = match.group("timestamp")
            ip = match.group("ip")
            timestamp = parse_timestamp(timestamp_str)
            failed_attempts[ip].append(timestamp)

brute_force_ips = []
for ip, timestamps in failed_attempts.items():
    timestamps.sort()  
    for i in range(len(timestamps) - attempt_threshold + 1):
        if timestamps[i + attempt_threshold - 1] - timestamps[i] <= time_window:
            brute_force_ips.append(ip)
            break

print("Potential brute-force IPs:")
for ip in set(brute_force_ips):
    print(ip)
