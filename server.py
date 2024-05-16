import json
import win32pipe
import win32file

pipe_name = r'\\.\pipe\my_pipe'

# Create a named pipe server
server = win32pipe.CreateNamedPipe(
    pipe_name,
    win32pipe.PIPE_ACCESS_DUPLEX,
    win32pipe.PIPE_TYPE_MESSAGE | win32pipe.PIPE_READMODE_MESSAGE | win32pipe.PIPE_WAIT,
    win32pipe.PIPE_UNLIMITED_INSTANCES,
    65536,
    65536,
    0,
    None
)

print("Waiting for client to connect...")

# Wait for a client to connect
win32pipe.ConnectNamedPipe(server, None)

print("Client connected.")

# Read data from the client
client_data = win32file.ReadFile(server, 4096)[1]
decoded_client_data = client_data.decode('UTF-8')
print("Data from client:", decoded_client_data)

# Process received data
# (Here you can perform any operations with the received data)

# Prepare response data
response_data = "Data received by server: " + decoded_client_data
encoded_response = response_data.encode('UTF-8')

# Send response to the client
win32file.WriteFile(server, encoded_response)

# Close the named pipe server
win32file.CloseHandle(server)
