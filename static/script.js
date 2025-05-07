let answers = {};

fetch("/quiz")
  .then(res => res.json())
  .then(quiz => {
    const container = document.getElementById("quiz-container");
    quiz.forEach((q, index) => {
      const div = document.createElement("div");
      div.className = "question";
      div.innerHTML = `<p><strong>Q${index + 1}:</strong> ${q.question}</p>`;
      q.options.forEach(opt => {
        div.innerHTML += `
          <div class="options">
            <label>
              <input type="radio" name="q${q.id}" value="${opt}"/> ${opt}
            </label>
          </div>
        `;
      });
      container.appendChild(div);
    });
  });

document.getElementById("submit").addEventListener("click", () => {
  const results = document.querySelectorAll("input[type='radio']:checked");
  if (results.length < 4) {
    alert("Please answer all questions.");
    return;
  }
  document.getElementById("result").innerText = "Quiz submitted! (Scoring not yet implemented)";
});
