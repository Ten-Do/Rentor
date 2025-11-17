import { type ButtonHTMLAttributes, type ElementType, forwardRef } from 'react'

type Size = 'sm' | 'md' | 'lg'
type ButtonVariant = 'solid' | 'outline' | 'ghost'
type ColorScheme = 'primary' | 'neutral'

interface ButtonProps
  extends Omit<ButtonHTMLAttributes<HTMLButtonElement>, 'size'> {
  as?: ElementType
  size?: Size
  variant?: ButtonVariant
  colorScheme?: ColorScheme
  href?: string
}

const baseClasses =
  'transition-colors duration-200 flex items-center justify-center'

const sizeClasses: Record<Size, string> = {
  sm: 'h-8 px-3 text-sm rounded-md',
  md: 'h-10 px-4 text-base rounded-lg',
  lg: 'h-12 px-6 text-lg rounded-lg',
}

const variantClasses: Record<ButtonVariant, Record<ColorScheme, string>> = {
  solid: {
    primary:
      'bg-indigo-600 text-white hover:bg-indigo-500 active:bg-indigo-700 disabled:bg-indigo-800 disabled:opacity-50',
    neutral:
      'bg-gray-600 text-white hover:bg-gray-500 active:bg-gray-700 disabled:bg-gray-800 disabled:opacity-50',
  },
  outline: {
    primary:
      'border-2 border-indigo-600 text-indigo-400 bg-transparent hover:bg-indigo-600/10 active:bg-indigo-600/20 disabled:border-indigo-800 disabled:text-indigo-800 disabled:opacity-50',
    neutral:
      'border-2 border-gray-600 text-gray-400 bg-transparent hover:bg-gray-600/10 active:bg-gray-600/20 disabled:border-gray-800 disabled:text-gray-800 disabled:opacity-50',
  },
  ghost: {
    primary:
      'text-indigo-400 hover:bg-indigo-600/10 active:bg-indigo-600/20 disabled:text-indigo-800 disabled:opacity-50',
    neutral:
      'text-gray-400 hover:bg-gray-600/10 active:bg-gray-600/20 disabled:text-gray-800 disabled:opacity-50',
  },
}

export const Button = forwardRef<HTMLElement, ButtonProps>(
  (
    {
      size = 'md',
      variant = 'solid',
      colorScheme = 'primary',
      className = '',
      as,
      ...props
    },
    ref
  ) => {
    const classes = `${baseClasses} ${sizeClasses[size]} ${variantClasses[variant][colorScheme]} ${className}`
    const Component = as || 'button'

    return <Component ref={ref} className={classes} {...props} />
  }
)

Button.displayName = 'Button'
