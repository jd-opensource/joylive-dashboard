/**
 * 全局服务空间记忆工具
 * 用于在整个应用中共享和记忆用户选择的服务空间
 */

const GLOBAL_SPACE_CODE_KEY = 'global_space_code'

/**
 * 获取当前选中的服务空间代码
 * @returns {string|null} 服务空间代码
 */
export function getCurrentSpaceCode() {
    return localStorage.getItem(GLOBAL_SPACE_CODE_KEY)
}

/**
 * 设置当前选中的服务空间代码
 * @param {string} spaceCode - 服务空间代码
 */
export function setCurrentSpaceCode(spaceCode) {
    if (spaceCode) {
        localStorage.setItem(GLOBAL_SPACE_CODE_KEY, spaceCode)
    }
}

/**
 * 清除当前选中的服务空间代码
 */
export function clearCurrentSpaceCode() {
    localStorage.removeItem(GLOBAL_SPACE_CODE_KEY)
}

/**
 * 初始化服务空间选择
 * @param {Array} spaceOptions - 服务空间选项列表
 * @returns {string|undefined} 初始化后的服务空间代码
 */
export function initSpaceCode(spaceOptions) {
    if (!spaceOptions || spaceOptions.length === 0) {
        return undefined
    }

    const saved = getCurrentSpaceCode()
    const found = saved && spaceOptions.some((item) => item.code === saved)
    const spaceCode = found ? saved : spaceOptions[0].code

    setCurrentSpaceCode(spaceCode)
    return spaceCode
}
