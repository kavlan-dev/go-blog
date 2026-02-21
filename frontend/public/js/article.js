document.addEventListener("DOMContentLoaded", async () => {
  // Get article ID from URL
  const urlParams = new URLSearchParams(window.location.search);
  const articleId = urlParams.get("id");

  if (!articleId) {
    document.getElementById("article").innerHTML = "<p>ID не указано</p>";
    return;
  }

  try {
    const response = await fetch(`/api/articles/${articleId}`);
    if (!response.ok) {
      throw new Error("Ошибка при получении статьи");
    }

    const data = await response.json();
    const article = data.article;
    const contentDiv = document.getElementById("article");

    if (!article) {
      contentDiv.innerHTML = "<p>Статья не найдена.</p>";
      return;
    }

    // Format the date
    const date = new Date(
      article.created_at || article.updated_at || Date.now(),
    );
    const formattedDate = date.toLocaleDateString("ru-RU", {
      year: "numeric",
      month: "long",
      day: "numeric",
    });

    contentDiv.innerHTML = `
            <article>
                <h2>${article.title}</h2>
                <div class="article-date">Опубликовано ${formattedDate}</div>
                <div class="article-content">
                    ${article.content.replace(/\n/g, "<br>")}
                </div>
                <div class="article-actions">
                    <a href="/" class="read-more">Вернуться на главную</a>
                </div>
            </article>
        `;
  } catch (error) {
    console.error("Ошибка загрузки статьи:", error);
    const contentDiv = document.getElementById("article");
    contentDiv.innerHTML =
      "<p>Ошибка загрузки статьи. Пожалуйста, попробуйте позже.</p>";
  }
});
