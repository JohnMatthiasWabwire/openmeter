/**
 * Checks whether a given value is in PascalCase
 * @param value the value to check
 * @returns true if the value is in PascalCase
 */
export function isPascalCaseNoAcronyms(value) {
  if (value === undefined || value === null || value === '') {
    return true
  }

  return /^([A-Z][a-z0-9]+)*[A-Z]?$/.test(value)
}

/**
 * Checks whether a given value is in camelCase
 * @param value the value to check
 * @returns true if the value is in camelCase
 */
export function isCamelCaseNoAcronyms(value) {
  if (value === undefined || value === null || value === '') {
    return true
  }

  return /^[^a-zA-Z0-9]?[a-z][a-z0-9]*([A-Z][a-z0-9]+)*[A-Z]?$/.test(value)
}

/**
 * Checks whether a given value is in snake_case
 * @param value the value to check
 * @returns true if the value is in snake_case
 */
export function isSnakeCase(value) {
  if (value === undefined || value === null || value === '') {
    return true
  }

  return /^([a-z0-9]+_)*[a-z0-9]+$/.test(value)
}
