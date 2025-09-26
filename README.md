# Contact Management Application

Simple command-line contact management system built with Go:
- Modular design with separate packages
- Struct-based data modeling
- Function-oriented programming
- Clear separation of concerns

## Features

- Add new contacts with name, phone number, email, and address
- View all saved contacts
- Update existing contact information
- Delete contacts
- Interactive menu interface
- Input validation and error handling

## Project Structure

```
Tugas_1/
├── input/
│   └── input.go      # User input handling
├── service/
│   └── crud.go       # CRUD operations
├── structure/
│   └── contact.go    # Contact data structure
└── main.go          # Main application
```

## How to Run

1. Make sure you have Go installed on your system
2. Clone this repository
3. Navigate to the project directory
4. Run the application:
   ```
   go run main.go
   ```

## Usage

The application provides a simple menu interface:
1. Add Contact - Create new contact entries
2. View Contacts - Display all saved contacts
3. Update Contact - Modify existing contact information
4. Delete Contact - Remove contacts from the list
q. Exit - Close the application

## Input Controls
- Use 'q' to return to the main menu from any input prompt
- Invalid inputs are handled with appropriate error messages