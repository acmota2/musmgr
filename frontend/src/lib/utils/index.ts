export function createPreviewUrl(file: File) {
  return URL.createObjectURL(file);
}

export function checkStringFormField(formField: FormDataEntryValue | null, key: string): string {
  if (typeof formField === "string") {
    return formField;
  }

  throw new Error(`Value for ${key} should be of type string`);
}

export function checkFileFormField(formField: FormDataEntryValue | null, key: string): File | null {
  if (formField === null || formField instanceof File) {
    return formField;
  }

  throw new Error(`Value for ${key} should be null or File`);
}

export function capitalize(str: string): string {
  const strs = str.split(/[ _]/);
  return strs.map((s) => s.charAt(0).toUpperCase() + s.slice(1)).join(" ");
}
