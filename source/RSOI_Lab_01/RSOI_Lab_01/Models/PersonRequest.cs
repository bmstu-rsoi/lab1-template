using RSOI_Lab_01.Entities;
using System.ComponentModel.DataAnnotations;

namespace RSOI_Lab_01.Models
{
    public class PersonRequest
    {
        /// <summary>
        /// Имя юзера
        /// </summary>
        [Required]
        public string Name { get; set; }
        /// <summary>
        /// Возраст юзера
        /// </summary>
        public int? Age { get; set; }
        /// <summary>
        /// Адрес юзера
        /// </summary>
        public string? Adress { get; set; }
        /// <summary>
        /// Работа юзера
        /// </summary>
        public string? Work { get; set; }

        public PersonRequest(string name, int? age = null, string? adress = null, string? work = null)
        {
            Name = name;
            Age = age;
            Adress = adress;
            Work = work;
        }

        public Person ToPerson() => new(Name, Age, Adress, Work);
    }
}
