import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'
import api from '../lib/api'
import type { Post } from '../types'
import ProfileSection from '../components/ProfileSection'
import PostFeed from '../components/PostFeed'
import PostForm from '../components/PostForm'

export default function DashboardPage() {
  const { logout } = useAuth()
  const navigate = useNavigate()

  const [posts, setPosts] = useState<Post[]>([])
  const [postsLoading, setPostsLoading] = useState(true)
  const [postsError, setPostsError] = useState<string | null>(null)

  useEffect(() => {
    api
      .get<Post[]>('/posts')
      .then((res) => setPosts(Array.isArray(res.data) ? res.data : []))
      .catch((err) => {
        if (err.response?.status === 401) {
          logout()
          navigate('/login')
        } else {
          setPostsError('Could not load posts.')
        }
      })
      .finally(() => setPostsLoading(false))
  }, [logout, navigate])

  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  const handlePostCreated = (post: Post) => {
    setPosts((prev) => [post, ...prev])
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Navbar */}
      <nav className="bg-white border-b border-gray-200 px-6 py-4 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <div className="w-8 h-8 rounded-full bg-blue-600 flex items-center justify-center">
            <span className="text-white text-xs font-bold">D</span>
          </div>
          <span className="font-semibold text-gray-900">Doctor Platform</span>
        </div>
        <button
          onClick={handleLogout}
          className="text-sm text-gray-500 hover:text-gray-900 transition-colors"
        >
          Sign out
        </button>
      </nav>

      {/* Two-column layout */}
      <main className="max-w-6xl mx-auto px-6 py-8 grid grid-cols-1 md:grid-cols-[320px_1fr] gap-6">
        {/* Left column: profile + post form */}
        <aside className="space-y-6">
          <ProfileSection />
          <PostForm onPostCreated={handlePostCreated} />
        </aside>

        {/* Right column: post feed */}
        <section>
          <h2 className="text-base font-semibold text-gray-900 mb-4">Community Feed</h2>
          <PostFeed posts={posts} loading={postsLoading} error={postsError} />
        </section>
      </main>
    </div>
  )
}
