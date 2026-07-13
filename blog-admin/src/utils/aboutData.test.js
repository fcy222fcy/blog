import test from 'node:test'
import assert from 'node:assert/strict'
import { parseProjects, parseSkills } from './aboutData.js'

test('parses legacy wrapped skills into editable tag names', () => {
  const skills = parseSkills('{"skills":[{"name":"Go","level":90},{"name":"Vue","level":70}]}')

  assert.deepEqual(skills, ['Go', 'Vue'])
})

test('parses legacy wrapped projects into an editable project list', () => {
  const projects = parseProjects('{"projects":[{"name":"Gin Blog","url":"https://example.com"}]}')

  assert.deepEqual(projects, [{
    name: 'Gin Blog',
    description: '',
    url: 'https://example.com',
    icon: ''
  }])
})
