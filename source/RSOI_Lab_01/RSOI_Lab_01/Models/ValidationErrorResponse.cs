namespace RSOI_Lab_01.Models
{
    public class ValidationErrorResponse
    {
        /// <summary>
        /// Сообщение об ошибке
        /// </summary>
        public string? Message { get; set; }

        /// <summary>
        /// Ошибки
        /// </summary>
        public Dictionary<string, string>? Errors { get; set; }
    }
}
