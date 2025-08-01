import React, { useState } from 'react'
import { PhoneInput } from './PhoneInput'

export const PhoneInputDemo: React.FC = () => {
  const [phoneValue, setPhoneValue] = useState('')
  const [isValid, setIsValid] = useState(false)

  return (
    <div style={{ padding: '20px', maxWidth: '400px' }}>
      <h2>PhoneInput Component Demo</h2>
      
      <PhoneInput
        value={phoneValue}
        onChange={setPhoneValue}
        onValidationChange={setIsValid}
      />
      
      <div style={{ marginTop: '16px', fontSize: '14px', color: '#666' }}>
        <p>Current value: {phoneValue}</p>
        <p>Is valid: {isValid ? 'Yes' : 'No'}</p>
      </div>
      
      <div style={{ marginTop: '16px' }}>
        <h3>Test Cases:</h3>
        <button onClick={() => setPhoneValue('13800138000')}>
          Set Valid Phone (13800138000)
        </button>
        <br />
        <button onClick={() => setPhoneValue('123')} style={{ marginTop: '8px' }}>
          Set Invalid Phone (123)
        </button>
        <br />
        <button onClick={() => setPhoneValue('')} style={{ marginTop: '8px' }}>
          Clear
        </button>
      </div>
      
      <div style={{ marginTop: '16px' }}>
        <h3>Disabled State:</h3>
        <PhoneInput
          value={phoneValue}
          onChange={setPhoneValue}
          onValidationChange={setIsValid}
          disabled
          placeholder="Disabled input"
        />
      </div>
    </div>
  )
}

export default PhoneInputDemo