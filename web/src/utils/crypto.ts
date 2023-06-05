const CryptoSecret = '__SecretKey__'

export function encrypt(data: any, key = CryptoSecret): string {
  const text = JSON.stringify(data)
  let result = ''
  for (let i = 0; i < text.length; i++)
    result += String.fromCharCode(text.charCodeAt(i) ^ key.charCodeAt(i % key.length))
  return window.btoa(result)
}

export function decrypt(cipherText: string, key = CryptoSecret): any {
  cipherText = window.atob(cipherText)
  let result = ''
  for (let i = 0; i < cipherText.length; i++)
    result += String.fromCharCode(cipherText.charCodeAt(i) ^ key.charCodeAt(i % key.length))
  if (result)
    return JSON.parse(result)
  return null
}
