import json
import win32pipe
import win32file

# Define data to send
command = "your_command_here"  # Define your command here
n1 = 1 + 2j  # Example complex number
n2 = 3 + 4j  # Example complex number

pipe_name = r'\\.\pipe\my_pipe'

data = dict(command=command, a=[n1.real, n1.imag], b=[n2.real, n2.imag])
encoded = json.dumps(data).encode('UTF-8')

# Connect to the named pipe server
client = win32file.CreateFile(
    pipe_name,
    win32file.GENERIC_READ | win32file.GENERIC_WRITE,
    0,
    None,
    win32file.OPEN_EXISTING,
    0,
    None
)

# Send data to the server
win32file.WriteFile(client, encoded)

# Read response from the server
response_data = win32file.ReadFile(client, 4096)[1]
decoded_response = response_data.decode('UTF-8')
print("Response from server:", decoded_response)

# Close the connection
client.close()
