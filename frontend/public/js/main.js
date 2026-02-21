document.addEventListener("DOMContentLoaded", async () => {
  try {
    const response = await fetch("/api/articles");
    if (!response.ok) {
      throw new Error("Ошибка при загрузке статей");
    }

    const data = await response.json();
    const articles = data.articles || [];
    const contentDiv = document.getElementById("articles");

    if (articles.length === 0) {
      contentDiv.innerHTML = "<p>Статей не найдено.</p>";
      return;
    }

    contentDiv.innerHTML = "";
    articles.forEach((article) => {
      const articleCard = document.createElement("div");
      articleCard.className = "article-card";

      // Format the date
      const date = new Date(
        article.created_at || article.updated_at || Date.now(),
      );
      const formattedDate = date.toLocaleDateString("ru-RU", {
        year: "numeric",
        month: "long",
        day: "numeric",
      });

      articleCard.innerHTML = `
                <h2>${article.title}</h2>
                <div class="article-date">Опубликовано ${formattedDate}</div>
                <div class="article-excerpt">
                    ${article.content.substring(0, 200)}${article.content.length > 200 ? "..." : ""}
                </div>
                <a href="/article.html?id=${article.ID}" class="read-more">Читать далее</a>
            `;

      contentDiv.appendChild(articleCard);
    });
  } catch (error) {
    console.error("Ошибка при загрузке статей:", error);
    const contentDiv = document.getElementById("content");
    contentDiv.innerHTML =
      "<p>Ошибка загрузки статей. Пожалуйста, попробуйте позже.</p>";
  }
});
