describe('登录页面跳转修复测试', () => {
  beforeEach(() => {
    cy.clearLocalStorage()
    
    // Mock API responses
    cy.intercept('POST', '/api/v1/doctor/SendSms', {
      statusCode: 200,
      body: { Message: '验证码发送成功', Code: 200 }
    }).as('sendSms')
    
    cy.intercept('POST', '/api/v1/doctor/LoginDoctor', {
      statusCode: 200,
      body: { Message: '登录成功', Code: 200, DId: 123 }
    }).as('login')
  })

  it('应该在登录成功后正确跳转到首页', () => {
    cy.visit('/login')
    
    // 验证在登录页面
    cy.url().should('include', '/login')
    cy.contains('手机验证码登录').should('be.visible')
    
    // 填写手机号
    cy.get('input[type="tel"]').type('13800138000')
    
    // 发送验证码
    cy.contains('获取验证码').click()
    cy.wait('@sendSms')
    
    // 填写验证码
    cy.get('input[type="number"]').type('1234')
    
    // 点击登录按钮
    cy.contains('登录').click()
    cy.wait('@login')
    
    // 验证登录成功提示
    cy.contains('登录成功，正在跳转...').should('be.visible')
    
    // 验证跳转到首页
    cy.url().should('include', '/home', { timeout: 10000 })
    
    // 验证首页内容加载
    cy.get('.doctor-header').should('be.visible')
    cy.contains('医生').should('be.visible')
  })

  it('应该使用多重备用方案确保跳转成功', () => {
    // 模拟路由跳转可能失败的情况
    cy.window().then((win) => {
      // 保存原始的pushState方法
      const originalPushState = win.history.pushState
      let pushStateCallCount = 0
      
      // Mock pushState to fail first few times
      win.history.pushState = function(...args) {
        pushStateCallCount++
        if (pushStateCallCount <= 2) {
          throw new Error('Navigation failed')
        }
        return originalPushState.apply(this, args)
      }
    })
    
    cy.visit('/login')
    
    // 执行登录流程
    cy.get('input[type="tel"]').type('13800138000')
    cy.contains('获取验证码').click()
    cy.wait('@sendSms')
    cy.get('input[type="number"]').type('1234')
    cy.contains('登录').click()
    cy.wait('@login')
    
    // 即使路由跳转失败，也应该最终到达首页
    cy.url().should('include', '/home', { timeout: 15000 })
  })

  it('应该在已登录状态下自动跳转到首页', () => {
    // 设置已登录状态
    cy.window().then((win) => {
      win.localStorage.setItem('doctor_token', 'test_token')
      win.localStorage.setItem('doctor_info', JSON.stringify({
        DId: 123,
        Name: '测试医生',
        Phone: '13800138000',
        Email: '',
        Avatar: '',
        Status: 'active'
      }))
      win.localStorage.setItem('last_login_time', Date.now().toString())
    })
    
    // 访问登录页面
    cy.visit('/login')
    
    // 应该自动跳转到首页
    cy.url().should('include', '/home', { timeout: 5000 })
    cy.contains('测试医生').should('be.visible')
  })

  it('应该处理登录失败并保持在登录页面', () => {
    // Mock登录失败
    cy.intercept('POST', '/api/v1/doctor/LoginDoctor', {
      statusCode: 400,
      body: { Message: '验证码错误', Code: 400 }
    }).as('loginFail')
    
    cy.visit('/login')
    
    // 执行登录流程
    cy.get('input[type="tel"]').type('13800138000')
    cy.get('input[type="number"]').type('1234')
    cy.contains('登录').click()
    cy.wait('@loginFail')
    
    // 验证错误提示
    cy.contains('验证码错误').should('be.visible')
    
    // 验证仍在登录页面
    cy.url().should('include', '/login')
    
    // 验证可以重新尝试登录
    cy.get('input[type="number"]').should('have.value', '')
    cy.get('.login-button').should('not.be.disabled')
  })

  it('应该正确处理网络错误', () => {
    // Mock网络错误
    cy.intercept('POST', '/api/v1/doctor/LoginDoctor', { forceNetworkError: true }).as('networkError')
    
    cy.visit('/login')
    
    // 执行登录流程
    cy.get('input[type="tel"]').type('13800138000')
    cy.get('input[type="number"]').type('1234')
    cy.contains('登录').click()
    cy.wait('@networkError')
    
    // 验证网络错误提示
    cy.contains('网络连接失败').should('be.visible')
    
    // 验证仍在登录页面
    cy.url().should('include', '/login')
  })

  it('应该正确管理登录按钮状态', () => {
    // Mock慢速登录响应
    cy.intercept('POST', '/api/v1/doctor/LoginDoctor', {
      statusCode: 200,
      body: { Message: '登录成功', Code: 200, DId: 123 },
      delay: 2000
    }).as('slowLogin')
    
    cy.visit('/login')
    
    // 填写表单
    cy.get('input[type="tel"]').type('13800138000')
    cy.get('input[type="number"]').type('1234')
    
    // 点击登录
    cy.contains('登录').click()
    
    // 验证加载状态
    cy.contains('登录中...').should('be.visible')
    cy.get('.login-button').should('be.disabled')
    
    // 等待登录完成
    cy.wait('@slowLogin')
    
    // 验证状态恢复并跳转
    cy.url().should('include', '/home', { timeout: 5000 })
  })

  it('应该正确保存登录状态到本地存储', () => {
    cy.visit('/login')
    
    // 执行登录流程
    cy.get('input[type="tel"]').type('13800138000')
    cy.contains('获取验证码').click()
    cy.wait('@sendSms')
    cy.get('input[type="number"]').type('1234')
    cy.contains('登录').click()
    cy.wait('@login')
    
    // 验证本地存储
    cy.window().then((win) => {
      expect(win.localStorage.getItem('doctor_token')).to.equal('temp_token')
      
      const doctorInfo = JSON.parse(win.localStorage.getItem('doctor_info') || '{}')
      expect(doctorInfo.DId).to.equal(123)
      expect(doctorInfo.Phone).to.equal('13800138000')
      
      const lastLoginTime = win.localStorage.getItem('last_login_time')
      expect(lastLoginTime).to.not.be.null
      expect(parseInt(lastLoginTime || '0')).to.be.greaterThan(Date.now() - 10000)
    })
    
    // 验证跳转成功
    cy.url().should('include', '/home')
  })

  it('应该在token过期时重新登录', () => {
    // 设置过期的登录状态
    cy.window().then((win) => {
      win.localStorage.setItem('doctor_token', 'expired_token')
      win.localStorage.setItem('doctor_info', JSON.stringify({
        DId: 123,
        Name: '测试医生',
        Phone: '13800138000',
        Email: '',
        Avatar: '',
        Status: 'active'
      }))
      // 设置8天前的登录时间（超过7天过期时间）
      const expiredTime = Date.now() - 8 * 24 * 60 * 60 * 1000
      win.localStorage.setItem('last_login_time', expiredTime.toString())
    })
    
    // 访问登录页面
    cy.visit('/login')
    
    // 应该保持在登录页面（因为token已过期）
    cy.url().should('include', '/login')
    cy.contains('手机验证码登录').should('be.visible')
    
    // 验证过期状态被清理
    cy.window().then((win) => {
      expect(win.localStorage.getItem('doctor_token')).to.be.null
    })
  })
})