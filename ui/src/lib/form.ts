export function isInvalid(field: any) {
  return field.state.meta.isTouched && !field.state.meta.isValid;
}

export function getErrors(field: any): string[] {
  if (!field.state.meta.errors) return [];
  return field.state.meta.errors
    .map((e: any) => (typeof e === "string" ? e : e?.message || ""))
    .filter(Boolean);
}
