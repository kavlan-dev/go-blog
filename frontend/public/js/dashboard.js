document.addEventListener("DOMContentLoaded", async () => {
  // Check if user is authenticated
  const isAuthenticated = await Auth.checkAuth();
  if (!isAuthenticated) {
    window.location.href = "/login.html";
    return;
  }

  // Load articles
  await loadArticles();

  // Set up add article button
  const addArticleBtn = document.getElementById("add-article-btn");
  if (addArticleBtn) {
    addArticleBtn.addEventListener("click", () => {
      window.location.href = "/edit-article.html";
    });
  }
});

async function loadArticles() {
  try {
    const response = await Auth.fetchWithAuth("/api/articles");
    if (!response.ok) {
      throw new Error("Ошибка при загрузке статей");
    }

    const data = await response.json();
    const articles = data.articles || [];
    const articlesList = document.getElementById("articles-list");

    if (articles.length === 0) {
      articlesList.innerHTML = "<p>Записей нет.</p>";
      return;
    }

    articlesList.innerHTML = "";
    articles.forEach((article) => {
      const articleItem = document.createElement("div");
      articleItem.className = "article-item";

      // Format the date
      const date = new Date(
        article.created_at || article.updated_at || Date.now(),
      );
      const formattedDate = date.toLocaleDateString("ru-RU", {
        year: "numeric",
        month: "long",
        day: "numeric",
      });

      articleItem.innerHTML = `
                <div>
                    <h3>${article.title}</h3>
                    <div class="article-date">${formattedDate}</div>
                </div>
                <div class="article-actions">
                    <button class="edit-btn" data-id="${article.ID}">Редактировать</button>
                    <button class="delete-btn" data-id="${article.ID}">Удалить</button>
                </div>
            `;

      articlesList.appendChild(articleItem);
    });

    // Set up event listeners for edit and delete buttons
    document.querySelectorAll(".edit-btn").forEach((btn) => {
      btn.addEventListener("click", (e) => {
        const articleId = e.target.getAttribute("data-id");
        window.location.href = `/edit-article.html?id=${articleId}`;
      });
    });

    document.querySelectorAll(".delete-btn").forEach((btn) => {
      btn.addEventListener("click", async (e) => {
        const articleId = e.target.getAttribute("data-id");
        if (confirm("Вы уверены, что хотите удалить статью?")) {
          await deleteArticle(articleId);
        }
      });
    });
  } catch (error) {
    console.error("Ошибка загрузки статей:", error);
    const articlesList = document.getElementById("articles-list");
    articlesList.innerHTML =
      "<p>Ошибки загрузки статей. Пожалуйста, попробуйте позже.</p>";
  }
}

async function deleteArticle(articleId) {
  try {
    const response = await Auth.fetchWithAuth(
      `/api/admin/articles/${articleId}`,
      {
        method: "DELETE",
      },
    );

    if (!response.ok) {
      throw new Error("Ошибка удаления статьи");
    }

    // Refresh the articles list
    await loadArticles();
  } catch (error) {
    console.error("Ошибка удаления статьи:", error);
    alert("Ошибка удаления статьи: " + error.message);
  }
}
