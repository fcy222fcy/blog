const parseJson = (value) => {
  if (Array.isArray(value) || (value && typeof value === 'object')) return value
  if (typeof value !== 'string' || !value.trim()) return []

  try {
    return JSON.parse(value)
  } catch {
    return []
  }
}

export const parseSkills = (value) => {
  const parsed = parseJson(value)
  const items = Array.isArray(parsed) ? parsed : parsed?.skills

  if (!Array.isArray(items)) return []

  return items
    .map((skill) => (typeof skill === 'string' ? skill : skill?.name))
    .filter((name) => typeof name === 'string' && name.trim())
    .map((name) => name.trim())
}

export const parseProjects = (value) => {
  const parsed = parseJson(value)
  const items = Array.isArray(parsed) ? parsed : parsed?.projects

  if (!Array.isArray(items)) return []

  return items
    .filter((project) => project && typeof project === 'object')
    .map((project) => ({
      ...project,
      name: typeof project.name === 'string' ? project.name : '',
      description: typeof project.description === 'string' ? project.description : '',
      url: typeof project.url === 'string' ? project.url.trim() : '',
      icon: typeof project.icon === 'string' ? project.icon : ''
    }))
}
