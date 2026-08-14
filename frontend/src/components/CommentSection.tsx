import { useEffect, useState } from 'react'
import api from '../lib/api'
import type { Comment } from '../types'

interface Props {
  postId: number
}

function formatDate(iso: string): string {
  return new Date(iso).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

export default function CommentSection({ postId }: Props) {
  const [comments, setComments] = useState<Comment[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const [text, setText] = useState('')
  const [submitting, setSubmitting] = useState(false)
  const [submitError, setSubmitError] = useState<string | null>(null)
  const [validationError, setValidationError] = useState<string | null>(null)

  useEffect(() => {
    setLoading(true)
    setError(null)
    api
      .get<Comment[]>(`/posts/${postId}/comments`)
      .then((res) => setComments(Array.isArray(res.data) ? res.data : []))
      .catch(() => setError('Could not load comments.'))
      .finally(() => setLoading(false))
  }, [postId])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    const trimmed = text.trim()
    if (!trimmed) {
      setValidationError('Comment cannot be empty.')
      return
    }
    setValidationError(null)
    setSubmitting(true)
    setSubmitError(null)
    try {
      const res = await api.post<Comment>(`/posts/${postId}/comments`, {
        content: trimmed,
      })
      setComments((prev) => [...prev, res.data])
      setText('')
    } catch {
      setSubmitError('Failed to post comment. Please try again.')
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <div className="mt-4 border-t border-gray-100 pt-4 space-y-3">
      {loading ? (
        <div className="animate-pulse space-y-2">
          <div className="h-3 bg-gray-200 rounded w-3/4" />
          <div className="h-3 bg-gray-200 rounded w-1/2" />
        </div>
      ) : error ? (
        <p className="text-xs text-red-500">{error}</p>
      ) : comments.length === 0 ? (
        <p className="text-xs text-gray-400">No comments yet. Be the first!</p>
      ) : (
        <ul className="space-y-2">
          {comments.map((c) => (
            <li key={c.id} className="text-sm">
              <p className="text-gray-800">{c.content}</p>
              <p className="text-xs text-gray-400 mt-0.5">{formatDate(c.created_at)}</p>
            </li>
          ))}
        </ul>
      )}

      <form onSubmit={handleSubmit} className="space-y-2">
        <textarea
          value={text}
          onChange={(e) => {
            setText(e.target.value)
            if (validationError) setValidationError(null)
          }}
          maxLength={1000}
          rows={2}
          placeholder="Write a comment…"
          className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm resize-none focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        {validationError && (
          <p className="text-xs text-red-500">{validationError}</p>
        )}
        {submitError && (
          <p className="text-xs text-red-500">{submitError}</p>
        )}
        <button
          type="submit"
          disabled={submitting}
          className="px-4 py-1.5 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 disabled:opacity-50 transition-colors"
        >
          {submitting ? 'Posting…' : 'Post Comment'}
        </button>
      </form>
    </div>
  )
}
