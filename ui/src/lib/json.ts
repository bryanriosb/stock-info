/**
 * Transform JSON response to handle big integers as strings.
 * JavaScript loses precision with numbers > Number.MAX_SAFE_INTEGER (9007199254740991).
 * This converts large numeric IDs to strings before JSON.parse.
 */
export function transformBigIntIds(data: string): unknown {
  if (typeof data !== 'string') return data

  // Replace "id": followed by a number with 15+ digits with string version
  const transformed = data.replace(/"id"\s*:\s*(\d{15,})/g, '"id": "$1"')
  return JSON.parse(transformed)
}
