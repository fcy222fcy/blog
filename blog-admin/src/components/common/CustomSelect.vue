<template>
  <div class="custom-select-wrapper" ref="wrapperRef">
    <div class="custom-select" :class="{ open: open, 'has-value': hasValue }" @click="toggleOpen">
      <!-- 触发器显示内容 -->
      <template v-if="multiple">
        <div class="select-tags">
          <template v-if="selectedOptions.length">
            <span
              v-for="opt in selectedOptions"
              :key="String(opt.value)"
              class="select-tag-pill"
              @click.stop="toggleOption(opt.value)"
            >
              {{ opt.label }}
              <span class="select-tag-x" @click.stop="toggleOption(opt.value)">×</span>
            </span>
          </template>
          <span v-else class="select-placeholder">{{ placeholder || '请选择' }}</span>
        </div>
      </template>
      <template v-else>
        <span class="select-value" :class="{ placeholder: !selectedLabelValid }">
          {{ selectedLabelValid ? selectedLabel : (placeholder || '请选择') }}
        </span>
      </template>
      <span class="select-arrow" :class="{ rotated: open }">
        <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
      </span>
    </div>
    <div
      class="custom-dropdown"
      v-if="open"
      :style="{ minWidth: dropdownWidth + 'px', maxHeight: maxDropdownHeight + 'px' }"
      @click.stop
    >
      <!-- 多选时：搜索框 + 全选/清空 -->
      <div class="dropdown-toolbar" v-if="multiple && (searchable || showSelectAll)">
        <input
          v-if="searchable"
          v-model="searchKeyword"
          type="text"
          class="dropdown-search"
          :placeholder="searchPlaceholder || '搜索...'"
          @click.stop
        >
        <div class="dropdown-actions" v-if="showSelectAll && filteredOptions.length">
          <button type="button" class="da-btn" @click.stop="selectAll">{{ allChecked ? '清空' : '全选' }}</button>
        </div>
      </div>

      <div class="dropdown-list" :class="{ 'wrap-mode': multiple }">
        <div
          v-for="opt in filteredOptions"
          :key="String(opt.value)"
          class="dropdown-item"
          :class="{ active: isOptionActive(opt.value), disabled: opt.disabled }"
          @click.stop="selectOption(opt)"
        >
          <!-- 复选框（多选或强制 showCheckbox） -->
          <span
            v-if="multiple || showCheckbox"
            class="item-checkbox"
            :class="{ checked: isOptionActive(opt.value) }"
          >
            <svg v-if="isOptionActive(opt.value)" xmlns="http://www.w3.org/2000/svg" width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
          </span>
          <span class="item-label">{{ opt.label }}</span>
        </div>
      </div>
      <div v-if="filteredOptions.length === 0" class="dropdown-empty">无匹配项</div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount, nextTick } from 'vue'

const props = defineProps({
  modelValue: { type: [String, Number, Array], default: '' },
  options: { type: Array, default: () => [] },
  placeholder: { type: String, default: '' },
  multiple: { type: Boolean, default: false },
  showCheckbox: { type: Boolean, default: false },
  searchable: { type: Boolean, default: false },
  searchPlaceholder: { type: String, default: '' },
  showSelectAll: { type: Boolean, default: false },
  maxDropdownHeight: { type: Number, default: 260 },
})

const emit = defineEmits(['update:modelValue', 'change'])

const wrapperRef = ref(null)
const open = ref(false)
const dropdownWidth = ref(0)
const searchKeyword = ref('')

const isArr = (v) => Array.isArray(v)

const hasValue = computed(() => {
  if (props.multiple) return isArr(props.modelValue) && props.modelValue.length > 0
  return props.modelValue !== '' && props.modelValue !== null && props.modelValue !== undefined
})

const selectedLabel = computed(() => {
  if (props.multiple) return ''
  const opt = props.options.find(o => o.value === props.modelValue)
  return opt ? opt.label : ''
})
const selectedLabelValid = computed(() => !!selectedLabel.value)

const selectedOptions = computed(() => {
  if (!props.multiple || !isArr(props.modelValue)) return []
  const set = new Set(props.modelValue)
  return props.options.filter(o => set.has(o.value))
})

const filteredOptions = computed(() => {
  if (!props.searchable || !searchKeyword.value.trim()) return props.options
  const kw = searchKeyword.value.trim().toLowerCase()
  return props.options.filter(o => String(o.label ?? '').toLowerCase().includes(kw))
})

const allChecked = computed(() => {
  if (!props.multiple) return false
  if (!filteredOptions.value.length) return false
  const values = new Set(props.modelValue || [])
  return filteredOptions.value.every(o => values.has(o.value))
})

const isOptionActive = (value) => {
  if (props.multiple) return isArr(props.modelValue) && props.modelValue.includes(value)
  return props.modelValue === value
}

const emitUpdate = (next, meta = {}) => {
  emit('update:modelValue', next)
  emit('change', next, meta)
}

const toggleOption = (value) => {
  if (!props.multiple) return
  const cur = Array.isArray(props.modelValue) ? [...props.modelValue] : []
  const idx = cur.indexOf(value)
  if (idx >= 0) cur.splice(idx, 1)
  else cur.push(value)
  emitUpdate(cur, { source: 'toggle', value })
}

const selectOption = (opt) => {
  if (opt.disabled) return
  const value = opt.value
  if (props.multiple) {
    toggleOption(value)
    return
  }
  emitUpdate(value, { source: 'click', value })
  open.value = false
}

const selectAll = () => {
  if (!props.multiple) return
  const current = new Set(props.modelValue || [])
  const pool = filteredOptions.value.filter(o => !o.disabled)
  if (allChecked.value) {
    for (const o of pool) current.delete(o.value)
  } else {
    for (const o of pool) current.add(o.value)
  }
  emitUpdate([...current], { source: allChecked.value ? 'clear-all' : 'select-all' })
}

const toggleOpen = () => {
  open.value = !open.value
  if (open.value) {
    searchKeyword.value = ''
    nextTick(() => {
      if (wrapperRef.value) dropdownWidth.value = Math.max(wrapperRef.value.offsetWidth, 220)
    })
  }
}

const close = () => { open.value = false }

const handleOutside = (e) => {
  if (wrapperRef.value && !wrapperRef.value.contains(e.target)) close()
}

onMounted(() => document.addEventListener('click', handleOutside))
onBeforeUnmount(() => document.removeEventListener('click', handleOutside))

defineExpose({ close })
</script>

<style scoped>
.custom-select-wrapper {
  position: relative;
  width: 100%;
}

.custom-select {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  width: 100%;
  min-height: 44px;
  padding: 8px 12px;
  padding-right: 40px;
  border: 1px solid var(--card-separator-color);
  border-radius: 4px;
  background: var(--card-background);
  color: var(--card-text-color-main);
  cursor: pointer;
  transition: all 0.15s ease;
  position: relative;
}

.custom-select:hover {
  border-color: rgba(var(--accent-color-rgb), 0.45);
}

.custom-select.open,
.custom-select:focus-within {
  border-color: var(--accent-color);
  box-shadow: 0 0 0 3px rgba(var(--accent-color-rgb), 0.12);
}

.select-arrow {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: var(--card-text-color-tertiary);
  pointer-events: none;
  transition: transform 0.15s ease;
}
.select-arrow.rotated { transform: translateY(-50%) rotate(180deg); }

.select-value {
  font-size: 14px;
  color: var(--card-text-color-main);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  flex: 1;
}
.select-value.placeholder { color: var(--card-text-color-tertiary); }

.select-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  flex: 1;
  min-height: 22px;
}
.select-placeholder {
  font-size: 14px;
  color: var(--card-text-color-tertiary);
}

/* 触发器里已选标签（#前缀 + 4px 小圆角，风格同前台详情页） */
.select-tag-pill {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  min-height: 24px;
  padding: 0 8px 0 10px;
  font-size: 12px;
  font-weight: 600;
  line-height: 1;
  background: rgba(var(--accent-color-rgb), 0.1);
  color: var(--accent-color);
  border-radius: 4px;
  white-space: nowrap;
}
.select-tag-pill::before { content: '#'; opacity: 0.7; }
.select-tag-x {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 16px;
  height: 16px;
  font-size: 14px;
  line-height: 1;
  border-radius: 3px;
  opacity: 0.6;
  cursor: pointer;
}
.select-tag-x:hover { background: rgba(var(--accent-color-rgb), 0.18); opacity: 1; }

/* 下拉弹层 */
.custom-dropdown {
  position: absolute;
  top: calc(100% + 6px);
  left: 0;
  z-index: 1000;
  background: var(--card-background);
  border: 1px solid var(--card-separator-color);
  border-radius: 4px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
  padding: 6px 0;
  overflow-y: auto;
  animation: dropIn 0.12s ease-out;
}
@keyframes dropIn {
  from { opacity: 0; transform: translateY(-4px); }
  to { opacity: 1; transform: translateY(0); }
}

.dropdown-toolbar {
  padding: 0 10px 8px;
  border-bottom: 1px solid var(--card-separator-color);
  margin-bottom: 4px;
}
.dropdown-search {
  width: 100%;
  padding: 8px 10px;
  border: 1px solid var(--card-separator-color);
  border-radius: 4px;
  background: var(--body-background);
  color: var(--card-text-color-main);
  font-size: 13px;
  margin-bottom: 8px;
  box-sizing: border-box;
}
.dropdown-search:focus {
  outline: none;
  border-color: var(--accent-color);
  box-shadow: 0 0 0 3px rgba(var(--accent-color-rgb), 0.12);
}
.dropdown-actions { display: flex; justify-content: flex-end; }
.da-btn {
  padding: 4px 10px;
  font-size: 12px;
  font-weight: 500;
  background: rgba(var(--accent-color-rgb), 0.08);
  color: var(--accent-color);
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.15s;
}
.da-btn:hover { background: rgba(var(--accent-color-rgb), 0.16); }

.dropdown-list {
  display: flex;
  flex-direction: column;
}
.dropdown-list.wrap-mode {
  flex-direction: row;
  flex-wrap: wrap;
  gap: 6px;
  padding: 8px 10px;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 14px;
  font-size: 14px;
  color: var(--card-text-color-main);
  cursor: pointer;
  transition: background 0.12s ease;
  border-radius: 2px;
  margin: 0 4px;
}
.dropdown-list.wrap-mode .dropdown-item {
  margin: 0;
  padding: 6px 10px;
  border-radius: 6px;
  border: 1px solid var(--card-separator-color);
  background: var(--card-background);
  font-size: 13px;
  flex: 0 0 auto;
}
.dropdown-item:hover { background: var(--body-background); }
.dropdown-list.wrap-mode .dropdown-item:hover {
  border-color: rgba(var(--accent-color-rgb), 0.45);
  background: rgba(var(--accent-color-rgb), 0.04);
}
.dropdown-item.active {
  background: rgba(var(--accent-color-rgb), 0.1);
  color: var(--accent-color);
  font-weight: 600;
}
.dropdown-list.wrap-mode .dropdown-item.active {
  border-color: var(--accent-color);
  background: rgba(var(--accent-color-rgb), 0.1);
}
.dropdown-item.disabled {
  opacity: 0.45;
  cursor: not-allowed;
  background: transparent !important;
  color: var(--card-text-color-tertiary) !important;
  font-weight: 400 !important;
}

.item-checkbox {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
  border: 1.5px solid var(--card-text-color-tertiary);
  border-radius: 3px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition: all 0.12s ease;
  background: var(--card-background);
  color: transparent;
}
.item-checkbox.checked {
  border-color: var(--accent-color);
  background: var(--accent-color);
  color: #fff;
}

.item-label { flex: 1; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }

.dropdown-empty {
  padding: 12px;
  text-align: center;
  font-size: 13px;
  color: var(--card-text-color-tertiary);
}
</style>
