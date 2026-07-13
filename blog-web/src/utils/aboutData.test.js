import test from 'node:test'
import assert from 'node:assert/strict'
import { parseProjects, parseSkills } from './aboutData.js'

test('parses legacy skills for public display', () => {
  const skills = parseSkills('{"skills":[{"name":"Go","level":90}]}')

  assert.deepEqual(skills, ['Go'])
})

test('parses legacy projects with their external URL', () => {
  const projects = parseProjects('{"projects":[{"name":"DesktopSnap","url":"https://example.com"}]}')

  assert.equal(projects[0].url, 'https://example.com')
})
