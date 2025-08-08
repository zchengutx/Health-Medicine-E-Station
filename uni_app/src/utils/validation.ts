/**
 * 验证工具函数
 */

/**
 * 验证中国大陆手机号格式
 * @param phoneNumber 手机号码
 * @returns boolean 是否为有效的手机号格式
 */
export const validatePhoneNumber = (phoneNumber: string): boolean => {
  // 中国大陆手机号正则：1开头，第二位为3-9，总共11位数字
  const phoneRegex = /^1[3-9]\d{9}$/;
  return phoneRegex.test(phoneNumber);
};

/**
 * 验证验证码格式
 * @param code 验证码
 * @returns boolean 是否为有效的验证码格式
 */
export const validateVerificationCode = (code: string): boolean => {
  // 验证码通常为4-6位数字
  const codeRegex = /^\d{4,6}$/;
  return codeRegex.test(code);
};

/**
 * 验证是否为空字符串
 * @param value 待验证的值
 * @returns boolean 是否为空
 */
export const isEmpty = (value: string): boolean => {
  return !value || value.trim().length === 0;
};

/**
 * 验证字符串长度是否在指定范围内
 * @param value 待验证的字符串
 * @param min 最小长度
 * @param max 最大长度
 * @returns boolean 是否在范围内
 */
export const validateLength = (value: string, min: number, max: number): boolean => {
  const length = value.trim().length;
  return length >= min && length <= max;
};