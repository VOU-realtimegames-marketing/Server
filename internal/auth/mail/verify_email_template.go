package mail

const VerifyEmailTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>OTP Verification</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f9f9f9;
            margin: 0;
            padding: 0;
        }
        .email-container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #ffffff;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            overflow: hidden;
        }
        .email-header {
            background-color: #4CAF50;
            color: #ffffff;
            text-align: center;
            padding: 20px;
            font-size: 24px;
        }
        .email-body {
            padding: 20px;
            color: #333333;
            line-height: 1.6;
        }
        .otp-code {
            display: inline-block;
            font-size: 28px;
            font-weight: bold;
            color: #4CAF50;
            margin: 20px 0;
            padding: 10px 20px;
            background-color: #f0f0f0;
            border-radius: 4px;
        }
        .footer {
            text-align: center;
            font-size: 14px;
            color: #777777;
            padding: 20px;
            background-color: #f2f2f2;
        }
        .footer a {
            color: #4CAF50;
            text-decoration: none;
        }
    </style>
</head>
<body>
    <div class="email-container">
        <div class="email-header">
            Welcome to VOU Marketing Platform!
        </div>
        <div class="email-body">
            <p>Hello %s,</p>
            <p>Thank you for signing up with us! To complete your registration, please use the One-Time Password (OTP) below:</p>
            <div class="otp-code">%s</div>
            <p>This OTP is valid for the next 15 minutes. Please do not share this code with anyone.</p>
            <p>If you did not request this, please ignore this email or contact our support team.</p>
            <p>Best regards,</p>
            <p><strong>VOU Marketing Team</strong></p>
        </div>
        <div class="footer">
            <p>If you have any questions, please contact our support at <a href="mailto:voumarketingteam@gmail.com">voumarketingteam@gmail.com</a>.</p>
        </div>
    </div>
</body>
</html>
`
