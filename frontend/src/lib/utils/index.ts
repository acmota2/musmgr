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
  console.log("Actual form field: ", formField, typeof formField);
  if (formField === null || formField instanceof File) {
    return formField;
  }

  throw new Error(`Value for ${key} should be null or File`);
}
