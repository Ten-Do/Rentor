import { type SelectHTMLAttributes, forwardRef } from 'react'

type Size = 'sm' | 'md' | 'lg'
type SelectVariant = 'default' | 'error'
type ColorScheme = 'primary' | 'neutral'

interface SelectProps
  extends Omit<SelectHTMLAttributes<HTMLSelectElement>, 'size'> {
  size?: Size
  variant?: SelectVariant
  colorScheme?: ColorScheme
}

const sizeClasses: Record<Size, string> = {
  sm: 'h-8 px-3 pr-8 text-sm rounded-md',
  md: 'h-10 px-4 pr-10 text-base rounded-lg',
  lg: 'h-12 px-5 pr-12 text-lg rounded-lg',
}

const baseClasses =
  "w-full bg-gray-800 text-white border appearance-none transition-colors duration-200 focus:outline-none disabled:opacity-50 disabled:cursor-not-allowed bg-[url(\"data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%239CA3AF' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3e%3cpolyline points='6 9 12 15 18 9'%3e%3c/polyline%3e%3c/svg%3e\")] bg-no-repeat bg-right"

const variantClasses: Record<SelectVariant, Record<ColorScheme, string>> = {
  default: {
    primary: `${baseClasses} border-gray-600 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20`,
    neutral: `${baseClasses} border-gray-600 focus:border-gray-500 focus:ring-2 focus:ring-gray-500/20`,
  },
  error: {
    primary: `${baseClasses} border-red-500 focus:border-red-400 focus:ring-2 focus:ring-red-500/20`,
    neutral: `${baseClasses} border-red-500 focus:border-red-400 focus:ring-2 focus:ring-red-500/20`,
  },
}

export const Select = forwardRef<HTMLSelectElement, SelectProps>(
  (
    {
      size = 'md',
      variant = 'default',
      colorScheme = 'primary',
      className = '',
      ...props
    },
    ref
  ) => {
    const classes = `${sizeClasses[size]} ${variantClasses[variant][colorScheme]} ${className}`

    return <select ref={ref} className={classes} {...props} />
  }
)

Select.displayName = 'Select'
