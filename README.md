# fretboard-games

CLI-based game to help me memorise notes/scales on a fretboard

This is work in progress for 2025-2026. 

The goal is to create a command-line game that helps me memorise the notes and scales on a guitar fretboard.
It should be able to help me answer questions like:

| Questions                                         | Game Name  |     Implemented      |
|---------------------------------------------------|:-----------|:--------------------:|
| List all C notes on the fretboard in strings 1-4. | `findnote` | ✅ Implemented       |
| What note is on the 5th fret of the 2nd string?   |            | 📋 To be implemented |
| What are the notes in a C major scale             |            | 📋 To be implemented |
| What are the notes in a G major chord ?           |            | 📋 To be implemented |
| What are the notes of C mixolydian scale ?        |            | 📋 To be implemented |

## Demo

https://github.com/user-attachments/assets/e75632dc-4d9e-4e18-bc82-7c0046c31067

## Wishlist

- [ ] configurable to allow for different tunings
- [ ] configurable number of strings
- [X] High score tracking
- [ ] Capture time taken to answer questions
- [ ] Different game modes (e.g., timed mode, survival mode)

## How to run

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/PauloMigAlmeida/fretboard-games.git
   cd fretboard-games
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Build the application:**
   ```bash
   go build -o fretboard-games .
   ```

4. **Run the CLI utility:**
   ```bash
   # Show available commands
   ./fretboard-games --help
   
   # Play the findnote game
   ./fretboard-games findnote
   ```

## Contribution

Know a good game that could improve one's learning and understanding of the fretboard?
Say no more! PRs are welcome :) 
