# Groupie Tracker

**Groupie Tracker** is a web application that provides detailed information about musicians and bands, including names, their members, creation dates, first albums, locations of concerts, and dates of concerts. The application fetches data from an external API and presents it in a user-friendly interface.

## Features

- **Artist Overview**: Displays a list of artists with their images and names.   
   Click on an artist to view details
- **Artist Details**: Provides comprehensive information about each artist, including:
  - Members
  - Creation Date
  - First Album
  - Concert Locations
  - Concert Dates
- **Search Bar**: Provides search on artist dB and gives suggestions by field (ie. creation date, first album date etc)   
   Click on a suggestion or type your search and press enter or hit search button
- **Filters**: Provides a variety of filters"
  - By range of creation date
  - By range of first album date
  - Radio buttons for number of members
  - input box for tour location

## APIs

- All artists 'GET' `serverAddr/api/artists`
- Search 'GET' `serverAddr/api/search?q=query`
- Filters 'POST' `serverAddr/api/filter`
   ```
   - Creation   []int  `json:"creation"`
   - FirstAlbum []int  `json:"firstAlbum"`
   - Members    []int  `json:"members"`
   - Locations  string `json:"locations"`
   ```

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/vstefanopoulos/groupie-tracker
   cd groupie-tracker
   ```

2. **Install Dependencies**:
   Ensure you have [Go](https://golang.org/dl/) installed. Then, run:
   ```bash
   go mod tidy
   ```

3. **Run the Application**:
   ```bash
   make build
   ```
   ```bash
   make run
   ```
   Optionally you can run on network by adding your IP address and the selected port seperated by ':' as an argument
   ```bash
   make run port=<YourIP>:<PORT>
   ```
   Run in one go on `localhost:8080/`:
   ```bash
   make debug
   ```

4. **Access the Application**:
   Open your browser and navigate to `http://localhost:8080` to use the Groupie Tracker.
   Or go to IP address with selected port
5. **Remove Binary**
   ```bash
   make clean
   ```

## Usage

- **Homepage**: Landing static html page.
- **Artists Page**: View a list of artists. 
   - **Search**: Option to search throught artists by any field.
   - **Filters**: Option to filters artists by
      - Range of creation dates
      - Range of first album dates
      - Number of members
      - Tour locations
- **Artist Details**: Click on an artist to see detailed information.

## Project Structure

- `main.go`: Entry point of the application.
- `backend/`: Server side. Offers three APIs: AllArtists (GET), search (GET), filter (POST)
- `ui/`: Contains static assets like CSS, JavaScript, and images.
- `templates/`: HTML templates for rendering web pages.

## Test
   Some tests of APIs are included. 
   ```bash
   make test
   ```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## Credits

- Vagelis Stefanopoulos   
- Sotirios Masteas   
- Alexandros Zachos