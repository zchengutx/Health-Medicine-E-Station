/// <reference types="cypress" />

declare global {
  namespace Cypress {
    interface Chainable {
      mockMobileViewport(): Chainable<void>
      mockApiResponse(method: string, url: string, response: any): Chainable<void>
      loginAsDoctor(phone?: string, password?: string): Chainable<void>
      fillPhoneAndSms(phone: string, smsCode: string): Chainable<void>
    }
  }
}

// Custom command to login as doctor
Cypress.Commands.add('loginAsDoctor', (phone = '13812345678', smsCode = '1234') => {
  cy.visit('/login')
  cy.fillPhoneAndSms(phone, smsCode)
  cy.get('[data-testid="login-button"]').click()
  cy.url().should('not.include', '/login')
})

// Custom command to fill phone and SMS
Cypress.Commands.add('fillPhoneAndSms', (phone: string, smsCode: string) => {
  cy.get('input[placeholder="请输入手机号码"]').type(phone)
  cy.get('.code-button').click()
  cy.get('input[placeholder="请输入验证码"]').type(smsCode)
})

export {}