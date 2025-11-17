import { type ReactNode } from 'react'

interface ParagraphProps {
  children: ReactNode
}

interface HeadingProps {
  children: ReactNode
}

const Paragraph = ({ children }: ParagraphProps) => {
  return <p>{children}</p>
}

const Heading = ({ children }: HeadingProps) => {
  return (
    <h2 className="text-2xl font-semibold text-indigo-400 mt-8 mb-4">
      {children}
    </h2>
  )
}

export const Typo = {
  Paragraph,
  Heading,
} as const
