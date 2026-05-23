import { useState } from 'react'
import api from '../lib/api'
import type { Post } from '../types'

interface Props {
  onPostCreated: (post: Post) => void
}

export default function PostForm({ onPostCreated }: Props) {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [communityId, setCommunityId] = useState('')
  const [imageUrl, setImageUrl] = useState('')
  const [submitting, setSubmitting] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!title.trim() || !content.trim() || !communityId) return

    setSubmitting(true)
    setError(null)
    try {
      const res = await api.post<Post>('/posts', {
        title: title.trim(),
        content: content.trim(),
        community_id: Number(communityId),
        image_url: imageUrl.trim(),
      })
      // Clear fields only on success
      setTitle('')
      setContent('')
      setCommunityId('')
      setImageUrl('')
      onPostCreated(res.data)
    } catch {
      setError('Failed to create post. Please try again.')
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <div className="bg-white rounded-xl border border-gray-200 p-5">
      <h2 className="text-base font-semibold text-gray-900 mb-4">New Post</h2>
      <form onSubmit={handleSubmit} className="space-y-3">
        <div>
          <label className="block text-xs text-gray-500 mb-1">
            Title <span className="text-red-400">*</span>
          </label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
            className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Post title"
          />
        </div>

        <div>
          <label className="block text-xs text-gray-500 mb-1">
            Content <span className="text-red-400">*</span>
          </label>
          <textarea
            value={content}
            onChange={(e) => setContent(e.target.value)}
            required
            rows={4}
            className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm resize-none focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="What's on your mind?"
          />
        </div>

        <div>
          <label className="block text-xs text-gray-500 mb-1">
            Community ID <span className="text-red-400">*</span>
          </label>
          <input
            type="number"
            value={communityId}
            onChange={(e) => setCommunityId(e.target.value)}
            required
            min={1}
            className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="e.g. 1"
          />
        </div>

        <div>
          <label className="block text-xs text-gray-500 mb-1">Image URL (optional)</label>
          <input
            type="text"
            value={imageUrl}
            onChange={(e) => setImageUrl(e.target.value)}
            className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="https://…"
          />
        </div>

        {error && <p className="text-xs text-red-500">{error}</p>}

        <button
          type="submit"
          disabled={submitting}
          className="w-full py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 disabled:opacity-50 transition-colors"
        >
          {submitting ? 'Posting…' : 'Post'}
        </button>
      </form>
    </div>
  )
}
