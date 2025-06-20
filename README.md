# Project: Bead & Breakfast
## Description: 
A simple web application for managing a bed and breakfast.
Based on the Udemy course [Building modern web applications with Go](https://www.udemy.com/course/building-modern-web-applications-with-go/).
## Business Needs:
- Bookings & Reservations
- Room Management (2 rooms)
## Key Functionality:
- Manage bookings (one or more nights)
- Manage room availability
- Book a room
- Notification system (guests and owners)
- Admin panel for managing bookings
  - review bookings
  - cancel bookings
  - update bookings
  - show calendar of bookings
## Features:
- Authentication System (owner)
- Database
- Email/Text notification system
## Dependencies:
- Go Version: 1.24.2
- [Chi Router](https://github.com/go-chi/chi)
- [SCS session management](https://github.com/alexedwards/scs)
- [NoSurf](https://github.com/justinas/nosurf)

## Project Structure:
The template section in this project dose not follow the approach of the course. 
Instead, it uses a more modular approach with separate template files for each page and 
partial templates for frameworks like bootstrap and sweetalert as well as loading custom css files and navigation. 
This allows for better organization and easier maintenance of the templates.