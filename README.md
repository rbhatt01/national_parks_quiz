# 🌲 Which National Park Are You? Quiz

A fun, interactive personality quiz that matches you with one of America's 63 national parks based on your traits and preferences!

## ✨ Features

- 🎯 15 carefully crafted questions
- 🏞️ All 63 US National Parks with official NPS images
- 📊 Personality trait matching algorithm
- 🎨 Beautiful, responsive design with custom green theme
- 📱 Mobile-friendly interface
- 🔄 Share your results

## 🚀 Live Demo

[Add your deployed URL here]

## 🛠️ Tech Stack

- **Backend**: Go 1.25+
- **Frontend**: HTML/CSS with Tailwind CSS
- **Fonts**: Outfit (Google Fonts)
- **Images**: Official National Park Service API

## 📁 Project Structure

```
national_parks_quiz/
├── data/
│   ├── parks.json          # All 63 parks with traits and descriptions
│   └── questions.json      # Quiz questions and options
├── internal/
│   ├── handlers/          # HTTP route handlers
│   ├── middleware/        # Logging middleware
│   ├── models/           # Data models (Park, Question, etc.)
│   └── services/         # Business logic (scoring, data loading)
├── templates/
│   ├── base.html         # Base template with layout
│   ├── home.html         # Landing page
│   ├── quiz.html         # Question pages
│   └── results.html      # Results display
├── main.go               # Application entry point
└── go.mod               # Go dependencies

```

## 🏃 Running Locally

1. **Clone the repository**
   ```bash
   git clone https://github.com/YOUR_USERNAME/national_parks_quiz.git
   cd national_parks_quiz
   ```

2. **Run the application**
   ```bash
   go run main.go
   ```

3. **Open your browser**
   ```
   http://localhost:8080
   ```

## 📊 How It Works

The quiz uses a trait-matching algorithm that:
1. Collects answers across 15 questions
2. Maps answers to personality traits (energy, social, remoteness, grit, drama, etc.)
3. Calculates similarity scores with all 63 parks
4. Returns your best match with a personalized description

## 🎨 Design Philosophy

- **Fun & Casual**: Uses the Outfit font and emoji for a friendly vibe
- **Green Theme**: Park-inspired color palette
- **No Survey Vibes**: Playful language and engaging interactions
- **Mobile-First**: Responsive design that works everywhere

## 🌟 Acknowledgments

- Park data and images from the [National Park Service](https://www.nps.gov/)
- Built with ❤️ for nature lovers and adventure seekers

---

**Made with 🌲 by Ria Bhatt**
