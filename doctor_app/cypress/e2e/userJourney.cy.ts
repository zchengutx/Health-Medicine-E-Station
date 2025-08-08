describe('Doctor App User Journey', () => {
  beforeEach(() => {
    cy.mockMobileViewport()
    
    // Mock API responses
    cy.mockApiResponse('POST', '/api/v1/doctor/SendSms', {
      Message: '验证码已发送',
      Code: 200
    })
    
    cy.mockApiResponse('POST', '/api/v1/doctor/LoginDoctor', {
      Message: '登录成功',
      Code: 200,
      DId: 1
    })
    
    cy.mockApiResponse('POST', '/api/v1/doctor/RegisterDoctor', {
      Message: '注册成功',
      Code: 200
    })
  })

  describe('Splash Screen', () => {
    it('should show splash screen and auto-navigate', () => {
      cy.visit('/')
      
      // Should show splash screen
      cy.get('.splash-view').should('be.visible')
      cy.contains('优医').should('be.visible')
      cy.contains('您身边的健康管理专家').should('be.visible')
      
      // Should show countdown
      cy.get('.skip-button').should('be.visible')
      cy.contains('跳过').should('be.visible')
      
      // Should show illustration
      cy.get('.illustration').should('be.visible')
      cy.get('.main-illustration').should('be.visible')
    })

    it('should allow skipping splash screen', () => {
      cy.visit('/')
      
      cy.get('.skip-button').click()
      
      // Should navigate to login page (assuming user is not logged in)
      cy.url().should('include', '/login')
    })
  })

  describe('Login Flow', () => {
    it('should complete login flow successfully', () => {
      cy.visit('/login')
      
      // Should show login page
      cy.contains('手机验证码登录').should('be.visible')
      cy.contains('欢迎使用优医在线问诊医生版').should('be.visible')
      
      // Should show medical illustration
      cy.get('.illustration').should('be.visible')
      cy.get('.medicine-bottle').should('be.visible')
      
      // Fill phone number
      cy.get('input[placeholder="请输入手机号码"]').type('13812345678')
      
      // Get SMS code button should be enabled
      cy.get('.code-button').should('not.be.disabled')
      cy.get('.code-button').click()
      
      // Fill SMS code
      cy.get('input[placeholder="请输入验证码"]').type('1234')
      
      // Login button should be enabled
      cy.get('.login-button, [data-testid="login-button"]').should('not.be.disabled')
      cy.get('.login-button, [data-testid="login-button"]').click()
      
      // Should navigate to home page
      cy.url().should('include', '/home')
    })

    it('should show validation errors for invalid input', () => {
      cy.visit('/login')
      
      // Try invalid phone number
      cy.get('input[placeholder="请输入手机号码"]').type('123')
      cy.get('.code-button').should('be.disabled')
      
      // Clear and enter valid phone
      cy.get('input[placeholder="请输入手机号码"]').clear().type('13812345678')
      cy.get('.code-button').should('not.be.disabled')
      
      // Login button should be disabled without SMS code
      cy.get('.login-button, [data-testid="login-button"]').should('be.disabled')
    })

    it('should navigate to register page', () => {
      cy.visit('/login')
      
      cy.get('.register-link span').click()
      cy.url().should('include', '/register')
    })
  })

  describe('Register Flow', () => {
    it('should complete register flow successfully', () => {
      cy.visit('/register')
      
      // Should show register page
      cy.contains('注册医生账号').should('be.visible')
      cy.contains('欢迎使用优医在线问诊医生版').should('be.visible')
      
      // Fill phone number
      cy.get('input[placeholder="请输入手机号码"]').type('13812345678')
      
      // Get SMS code
      cy.get('.code-button').click()
      
      // Fill SMS code
      cy.get('input[placeholder="请输入验证码"]').type('1234')
      
      // Fill password
      cy.get('input[placeholder="请输入密码"]').type('Abc123456')
      
      // Should show password strength indicator
      cy.get('.password-strength').should('be.visible')
      
      // Register button should be enabled
      cy.get('.register-button, [data-testid="register-button"]').should('not.be.disabled')
      cy.get('.register-button, [data-testid="register-button"]').click()
      
      // Should navigate to login page after successful registration
      cy.url().should('include', '/login')
    })

    it('should toggle password visibility', () => {
      cy.visit('/register')
      
      const passwordInput = cy.get('input[placeholder="请输入密码"]')
      passwordInput.should('have.attr', 'type', 'password')
      
      cy.get('.password-toggle').click()
      passwordInput.should('have.attr', 'type', 'text')
      
      cy.get('.password-toggle').click()
      passwordInput.should('have.attr', 'type', 'password')
    })

    it('should show password strength feedback', () => {
      cy.visit('/register')
      
      const passwordInput = cy.get('input[placeholder="请输入密码"]')
      
      // Weak password
      passwordInput.type('123')
      cy.get('.password-strength').should('be.visible')
      cy.get('.strength-fill.weak').should('exist')
      
      // Clear and type stronger password
      passwordInput.clear().type('Abc123456!')
      cy.get('.strength-fill.strong').should('exist')
    })

    it('should navigate to login page', () => {
      cy.visit('/register')
      
      cy.get('.login-link span').click()
      cy.url().should('include', '/login')
    })
  })

  describe('Responsive Design', () => {
    it('should work on different screen sizes', () => {
      // Test on iPhone SE
      cy.viewport(320, 568)
      cy.visit('/login')
      cy.get('.login-view').should('be.visible')
      
      // Test on iPhone 12
      cy.viewport(390, 844)
      cy.visit('/login')
      cy.get('.login-view').should('be.visible')
      
      // Test on iPad
      cy.viewport(768, 1024)
      cy.visit('/login')
      cy.get('.login-view').should('be.visible')
    })
  })

  describe('Error Handling', () => {
    it('should handle API errors gracefully', () => {
      // Mock API error
      cy.mockApiResponse('POST', '/api/v1/doctor/SendSms', {
        statusCode: 400,
        body: {
          Message: '手机号格式错误',
          Code: 400
        }
      })
      
      cy.visit('/login')
      cy.get('input[placeholder="请输入手机号码"]').type('13812345678')
      cy.get('.code-button').click()
      
      // Should show error message (assuming toast is implemented)
      // This would depend on how error messages are displayed
    })

    it('should handle network errors', () => {
      // Mock network error
      cy.intercept('POST', '/api/v1/doctor/SendSms', { forceNetworkError: true })
      
      cy.visit('/login')
      cy.get('input[placeholder="请输入手机号码"]').type('13812345678')
      cy.get('.code-button').click()
      
      // Should handle network error gracefully
    })
  })
})