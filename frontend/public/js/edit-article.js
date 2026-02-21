document.addEventListener("DOMContentLoaded", async () => {
  // Check if user is authenticated
  const isAuthenticated = await Auth.checkAuth();
  if (!isAuthenticated) {
    window.location.href = "/login.html";
    return;
  }

  // Get article ID from URL
  const urlParams = new URLSearchParams(window.location.search);
  const articleId = urlParams.get("id");

  const titleInput = document.getElementById("title");
  const contentInput = document.getElementById("content");
  const saveBtn = document.getElementById("save-btn");
  const deleteBtn = document.getElementById("delete-btn");

  // If we have an article ID, load the article
  if (articleId) {
    try {
      const response = await Auth.fetchWithAuth(`/api/articles/${articleId}`);
      if (!response.ok) {
        throw new Error("Ошибка загрузки статьи");
      }

      const data = await response.json();
      const article = data.article;

      if (article) {
        titleInput.value = article.title;
        contentInput.value = article.content;
      }
    } catch (error) {
      console.error("Ошибка загрузки статьи:", error);
      alert("Ошибка загрузки статьи: " + error.message);
    }
  } else {
    // Hide delete button for new articles
    deleteBtn.style.display = "none";
  }

  // Set up form submission
  const articleForm = document.getElementById("article-form");
  if (articleForm) {
    articleForm.addEventListener("submit", async (e) => {
      e.preventDefault();

      const title = titleInput.value;
      const content = contentInput.value;

      try {
        let response;
        if (articleId) {
          // Update existing article
          response = await Auth.fetchWithAuth(
            `/api/admin/articles/${articleId}`,
            {
              method: "PUT",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({
                title: title,
                content: content,
              }),
            },
          );
        } else {
          // Create new article
          response = await Auth.fetchWithAuth("/api/admin/articles", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              title: title,
              content: content,
            }),
          });
        }

        if (!response.ok) {
          throw new Error("Ошибка сохранения статьи");
        }

        // Redirect to dashboard after successful save
        window.location.href = "/dashboard.html";
      } catch (error) {
        console.error("Ошибка сохранения статьи:", error);
        alert("Ошибка сохранения статьи: " + error.message);
      }
    });
  }

  // Set up delete button
  if (deleteBtn && articleId) {
    deleteBtn.addEventListener("click", async () => {
      if (confirm("Вы уверены, что хотите удалить эту статью?")) {
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

          // Redirect to dashboard after successful delete
          window.location.href = "/dashboard.html";
        } catch (error) {
          console.error("Ошибка удаления статьи:", error);
          alert("Ошибка удаления статьи: " + error.message);
        }
      }
    });
  }
});
