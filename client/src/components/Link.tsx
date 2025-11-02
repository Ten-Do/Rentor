import { type AnchorHTMLAttributes, forwardRef } from 'react'
import {
  Link as RouterLink,
  type LinkProps as RouterLinkProps,
} from 'react-router-dom'

type Size = 'sm' | 'md' | 'lg'
type LinkVariant = 'default' | 'underlined'
type ColorScheme = 'primary' | 'neutral'

interface LinkProps
  extends Omit<AnchorHTMLAttributes<HTMLAnchorElement>, 'href'> {
  to?: string
  href?: string
  size?: Size
  variant?: LinkVariant
  colorScheme?: ColorScheme
}

const sizeClasses: Record<Size, string> = {
  sm: 'text-sm',
  md: 'text-base',
  lg: 'text-lg',
}

const variantClasses: Record<LinkVariant, Record<ColorScheme, string>> = {
  default: {
    primary:
      'transition-colors duration-200 text-indigo-400 hover:text-indigo-300',
    neutral: 'transition-colors duration-200 text-gray-400 hover:text-gray-300',
  },
  underlined: {
    primary:
      'transition-colors duration-200 text-indigo-400 underline hover:text-indigo-300',
    neutral:
      'transition-colors duration-200 text-gray-400 underline hover:text-gray-300',
  },
}

export const Link = forwardRef<HTMLAnchorElement, LinkProps>(
  (
    {
      to,
      href,
      size = 'md',
      variant = 'default',
      colorScheme = 'primary',
      className = '',
      ...props
    },
    ref
  ) => {
    const classes = `${sizeClasses[size]} ${variantClasses[variant][colorScheme]} ${className}`

    if (to) {
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const { to: _, ...restProps } = props as RouterLinkProps & { to?: string }
      return <RouterLink ref={ref} to={to} className={classes} {...restProps} />
    }

    return <a ref={ref} href={href} className={classes} {...props} />
  }
)
