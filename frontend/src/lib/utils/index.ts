export function createPreviewUrl(file: File) {
  return URL.createObjectURL(file);
}

export function checkStringFormField(
  formField: FormDataEntryValue | null,
  key: string,
): string | never {
  console.log(formField);
  if (typeof formField === "string") {
    return formField;
  }

  throw new Error(`Value for ${key} should be of type string`);
}
