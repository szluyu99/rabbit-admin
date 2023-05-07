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

// import CryptoJS from 'crypto-js'

// export function encrypt(data: any): string {
// const newData = JSON.stringify(data)
// return CryptoJS.AES.encrypt(newData, CryptoSecret).toString()
// }

// export function decrypt(cipherText: string) {
// const bytes = CryptoJS.AES.decrypt(cipherText, CryptoSecret)
// const originalText = bytes.toString(CryptoJS.enc.Utf8)
// if (originalText)
//   return JSON.parse(originalText)
// return null
// }
