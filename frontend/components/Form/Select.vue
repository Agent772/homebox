<template>
  <div class="form-control w-full">
    <label class="label">
      <span class="label-text">{{ label }}</span>
    </label>
    <select class="select select-bordered" v-model="value">
      <option disabled selected>Pick one</option>
      <option v-for="obj in items" :value="obj">
        {{ name != '' ? obj[name] : obj }}
      </option>
    </select>
    <!-- <label class="label">
      <span class="label-text-alt">Alt label</span>
      <span class="label-text-alt">Alt label</span>
    </label> -->
  </div>
</template>

<script lang="ts" setup>
  const emit = defineEmits(['update:modelValue']);
  const props = defineProps({
    label: {
      type: String,
      default: '',
    },
    modelValue: {
      type: Object as any,
      default: null,
    },
    items: {
      type: Array as () => any[],
      required: true,
    },
    name: {
      type: String,
      default: 'name',
    },
    selectFirst: {
      type: Boolean,
      default: false,
    },
  });

  watchOnce(
    () => props.items,
    () => {
      if (props.selectFirst && props.items.length > 0) {
        value.value = props.items[0];
      }
    }
  );

  const value = useVModel(props, 'modelValue', emit);
</script>