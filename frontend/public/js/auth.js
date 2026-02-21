class Auth {
    static getToken() {
        return localStorage.getItem('token');
    }

    static setToken(token) {
        localStorage.setItem('token', token);
    }

    static removeToken() {
        localStorage.removeItem('token');
    }

    static isAuthenticated() {
        return !!this.getToken();
    }

    static async checkAuth() {
        const token = this.getToken();
        if (!token) {
            this.updateAuthUI(false);
            return false;
        }

        // Simple token validation - in production you might want to call an endpoint
        // to validate the token
        try {
            const payload = JSON.parse(atob(token.split('.')[1]));
            const expires = payload.exp * 1000;
            if (Date.now() >= expires) {
                this.removeToken();
                this.updateAuthUI(false);
                return false;
            }
            this.updateAuthUI(true);
            return true;
        } catch (e) {
            this.removeToken();
            this.updateAuthUI(false);
            return false;
        }
    }

    static updateAuthUI(isAuthenticated) {
        const loginLink = document.getElementById('login-link');
        const dashboardLink = document.getElementById('dashboard-link');
        const logoutLink = document.getElementById('logout-link');

        if (loginLink) {
            loginLink.style.display = isAuthenticated ? 'none' : 'inline';
        }
        if (dashboardLink) {
            dashboardLink.style.display = isAuthenticated ? 'inline' : 'none';
        }
        if (logoutLink) {
            logoutLink.style.display = isAuthenticated ? 'inline' : 'none';
        }

        if (logoutLink && isAuthenticated) {
            logoutLink.addEventListener('click', (e) => {
                e.preventDefault();
                this.removeToken();
                window.location.href = '/login.html';
            });
        }
    }

    static async fetchWithAuth(url, options = {}) {
        const token = this.getToken();
        if (token) {
            options.headers = {
                ...options.headers,
                'Authorization': `Bearer ${token}`
            };
        }

        const response = await fetch(url, options);

        if (response.status === 401) {
            this.removeToken();
            window.location.href = '/login.html';
        }

        return response;
    }
}

// Initialize auth check on page load
document.addEventListener('DOMContentLoaded', () => {
    Auth.checkAuth();
});
