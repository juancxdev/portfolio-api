package helper

import "fmt"

func GenerateEmailHTML(email string) (string, error) {
	return fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
    <head>
        <meta charset="UTF-8">
        <style>
            body {
                margin: 0;
                padding: 0;
                background-color: #ffffff;
            }
            .email-container {
                max-width: 600px;
                margin: 0 auto;
                padding: 40px 20px;
                font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
                color: #1a1a1a;
                text-align: center;
            }
            .logo {
                width: 120px;
                margin-bottom: 40px;
            }
            .title {
                font-size: 42px;
                font-weight: 700;
                margin: 0 0 20px 0;
                line-height: 1.2;
                color: #000000;
            }
            .subtitle {
                font-size: 24px;
                font-weight: normal;
                margin: 0 0 40px 0;
                line-height: 1.4;
                color: #4a4a4a;
            }
            .message {
                font-size: 18px;
                line-height: 1.6;
                margin: 0 0 30px 0;
                color: #666666;
            }
            .button {
                display: inline-block;
                padding: 15px 30px;
                background-color: #007bff;
                color: #ffffff !important;
                text-decoration: none;
                border-radius: 8px;
                font-size: 16px;
                font-weight: 600;
                margin: 30px 0;
                transition: background-color 0.3s ease;
            }
            .button:hover {
                background-color: #0056b3;
            }
            .details {
                font-size: 14px;
                color: #666666;
                margin-top: 40px;
                line-height: 1.6;
            }
        </style>
    </head>
    <body>
        <div class="email-container">
            <h1 class="title">¡Tenemos buenas noticias!</h1>
            
            <h2 class="subtitle">Alguien está interesado en contratarte</h2>
            
            <p class="message">
                Hola Juan Campos,<br><br>
                Nos complace informarte que %s está interesado en contactarte.
            </p>
        </div>
    </body>
    </html>
    `, email), nil
}
