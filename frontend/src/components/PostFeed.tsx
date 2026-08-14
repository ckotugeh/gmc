import { useState } from 'react'
import type { Post } from '../types'
import CommentSection from './CommentSection'

interface Props {
  posts: Post[]
  loading: boolean
  error: string | null
}

function formatDate(iso: string): string {
  return new Date(iso).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

function truncate(text: string, max: number): string {
  if (text.length <= max) return text
  return text.slice(0, max) + '…'
}

export default function PostFeed({ posts, loading, error }: Props) {
  const [expandedPostId, setExpandedPostId] = useState<number | null>(null)

  const safePosts = Array.isArray(posts) ? posts : []

  const handlePostClick = (id: number) => {
    setExpandedPostId((prev) => (prev === id ? null : id))
  }

  if (loading) {
    return (
      <div className="space-y-3">
        {[1, 2, 3].map((i) => (
          <div key={i} className="bg-white rounded-xl border border-gray-200 p-5 animate-pulse">
            <div className="h-4 bg-gray-200 rounded w-1/2 mb-2" />
            <div className="h-3 bg-gray-200 rounded w-full mb-1" />
            <div className="h-3 bg-gray-200 rounded w-3/4" />
          </div>
        ))}
      </div>
    )
  }

  if (error) {
    return (
      <div className="bg-white rounded-xl border border-gray-200 p-5">
        <p className="text-sm text-red-500">{error}</p>
      </div>
    )
  }

  if (safePosts.length === 0) {
    return (
      <div className="bg-white rounded-xl border border-gray-200 p-5">
        <p className="text-sm text-gray-400 text-center py-6">No posts yet.</p>
      </div>
    )
  }

  return (
    <div className="space-y-3">
      {safePosts.map((post) => {
        const isExpanded = expandedPostId === post.id
        return (
          <div
            key={post.id}
            className="bg-white rounded-xl border border-gray-200 p-5 cursor-pointer hover:border-gray-300 transition-colors"
            onClick={() => handlePostClick(post.id)}
          >
            <h3 className="text-sm font-semibold text-gray-900 mb-1">{post.title}</h3>
            <p className="text-sm text-gray-600 mb-2">
              {isExpanded ? post.content : truncate(post.content, 200)}
            </p>
            {post.image_url && (
              <img
                src={post.image_url}
                alt={post.title}
                className="rounded-lg max-h-64 object-cover mb-2"
                onClick={(e) => e.stopPropagation()}
              />
            )}
            <p className="text-xs text-gray-400">{formatDate(post.created_at)}</p>

            {isExpanded && (
              <div onClick={(e) => e.stopPropagation()}>
                <CommentSection postId={post.id} />
              </div>
            )}
          </div>
        )
      })}
    </div>
  )
}
