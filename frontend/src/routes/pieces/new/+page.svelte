<script lang="ts">
  import ConfirmDialog from "$lib/ui/form/ConfirmDialog.svelte";
  import FormButtonSet from "$lib/ui/form/FormButtonSet.svelte";
  import FormField from "$lib/ui/form/FormField.svelte";
  import Select from "$lib/ui/form/Select.svelte";
  import TextArea from "$lib/ui/form/TextArea.svelte";
  import TextInput from "$lib/ui/form/TextInput.svelte";
  import { capitalize } from "$lib/utils/index.js";

  function confirm() {
    confirmDialog?.showModal();
  }

  function cancel() {
    history.back();
  }

  let confirmDialog: HTMLDialogElement | null = $state(null);
  let { data } = $props();
</script>

<div class="container narrow">
  <FormField method="POST" action="?/create" id="create-piece">
    <label for="title">
      <h2>Title</h2>
    </label>
    <TextInput name="title" required />
    <label for="composed_at">
      <h2>Composed at</h2>
    </label>
    <TextInput name="composed_at" required placeholder="YYYY" />
    <label for="instrumentation">
      <h2>Instrumentation</h2>
    </label>
    <Select name="instrumentation">
      {#each data.instrumentationNames as name}
        <option value="{name}">{capitalize(name)}</option>
      {/each}
    </Select>
    <label for="description">
      <h2>Description</h2>
    </label>
    <TextArea name="description" required />

    <FormButtonSet {cancel} {confirm} />
  </FormField>
</div>

<ConfirmDialog bind:confirmDialog formId="create-piece" />

<style>
  h2 {
    font-size: 28px;
  }

  label > * {
    margin: 6px 0;
  }
</style>
